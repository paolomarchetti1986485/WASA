<template>
  <div>
    <h1>Home Page</h1>
    <button @click="logout">Logout</button>
    
    <div v-if="loading">Loading...</div>
    <div v-if="errormsg">{{ errormsg }}</div>

    <div class="search-bar">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Cerca un utente..."
        @keydown.enter="searchUser"
      />
      <button @click="searchUser">Cerca</button>
    </div>

    <div v-for="photo in photos" :key="photo.photoId" class="photo">
      <div class="photo-container">
        <img :src="photo.imageSrc" alt="photo" class="photo-img" @error="handleImageError" />
      </div>
      <div class="photo-info">
        <!-- Nome utente cliccabile per reindirizzare al profilo -->
        <p @click="goToProfile(photo.userId)" class="user-profile-link" style="cursor: pointer;">
          Uploaded by: {{ photo.username || 'Unknown User' }}
        </p>
        <p>Uploaded at: {{ new Date(photo.uploadDateTime).toLocaleString() }}</p>
        <p>Likes: {{ photo.likes.length }}</p>
        <p>Comments: {{ photo.comments.length }}</p>
      </div>
      <div class="photo-actions">
        <button @click="toggleLike(photo)">Like</button>
        <button @click="commentPhoto(photo.photoId)">Comment</button>
      </div>
      <div v-for="comment in photo.comments" :key="comment.commentId" class="comment">
        <p>{{ comment.commentText }}</p>
        <button v-if="comment.userId == userId" @click="deleteComment(photo.photoId, comment.commentId)">Delete</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      loading: false,
      errormsg: null,
      photos: [],
      userId: localStorage.getItem('token'),
      searchQuery: '',  // Variabile per immagazzinare il testo della ricerca

    };
  },
  methods: {
    goToProfile(userId) {
      this.$router.push({ name: 'NonLoggedProfile', params: { userId: userId } });
    },
    async fetchPhotos() {
      this.loading = true;
      try {
        let response = await axios.get(`/user/${this.userId}/stream`);
        this.photos = response.data.map(photo => ({
          ...photo,
          comments: photo.Comments || [],
          likes: photo.Likes || []
        }));

        // Carica le immagini per ogni foto
        for (let photo of this.photos) {
          this.loadImage(photo.photoId);
        }
      } catch (e) {
        console.error("Error fetching photos:", e);
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    async loadImage(photoId) {
      try {
        const response = await axios.get(`/photos/${photoId}/image`, { responseType: 'blob' });
        const url = URL.createObjectURL(response.data);
        const photo = this.photos.find(p => p.photoId === photoId);
        if (photo) {
          photo.imageSrc = url;
        }
      } catch (e) {
        console.error("Error loading image:", e);
      }
    },
    handleImageError(event) {
      event.target.style.display = 'none'; // Nasconde l'immagine se c'è un errore
    },
    async toggleLike(photo) {
      const hasLiked = photo.likes.some(like => like.userId == this.userId);
      try {
        if (hasLiked) {
          await axios.delete(`/photos/${photo.photoId}/likes/${this.userId}`);
        } else {
          await axios.put(`/photos/${photo.photoId}/likes/${this.userId}`);
        }
        this.fetchPhotos();
      } catch (e) {
        console.error("Error toggling like on photo:", e);
        this.errormsg = e.toString();
      }
    },
    async commentPhoto(photoId) {
      const text = prompt("Enter your comment:");
      if (!text) return;
      try {
        await axios.post(`/user/${this.userId}/photos/${photoId}/comments/`, { text });
        this.fetchPhotos();
      } catch (e) {
        console.error("Error commenting on photo:", e);
        this.errormsg = e.toString();
      }
    },
    async deleteComment(photoId, commentId) {
      try {
        await axios.delete(`/user/${this.userId}/photos/${photoId}/comments/${commentId}`);
        this.fetchPhotos();
      } catch (e) {
        console.error("Error deleting comment:", e);
        this.errormsg = e.toString();
      }
    },
    logout() {
      localStorage.removeItem('token');
      this.$router.replace('/login');
    },
    searchUser() {
      if (this.searchQuery) {
        // Reindirizza alla view della ricerca passando il nome utente come parametro
        this.$router.push(`/search?username=${this.searchQuery}`);
      }
    }
  },
  mounted() {
    this.fetchPhotos();
  }
}
</script>

<style>
.photo {
  margin: 20px 0;
}

.photo-container {
  width: 300px; /* Dimensione fissa per larghezza */
  height: 300px; /* Dimensione fissa per altezza */
  border: 2px solid black; /* Bordo nero */
  display: flex;
  justify-content: center;
  align-items: center;
}
.photo-img {
  max-width: 100%;
  max-height: 100%;
}
</style>