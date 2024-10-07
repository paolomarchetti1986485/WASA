<template>
  <div>
    <h1 v-if="username">{{ username }}'s Profile Page</h1>
    <h1 v-else>Loading...</h1>
    <button @click="logout">Logout</button>
    <div>
      <button @click="editingUsername = !editingUsername">
        {{ editingUsername ? 'Cancel' : 'Update Username' }}
      </button>
      <div v-if="editingUsername">
        <input v-model="newUsername" placeholder="Enter new username" />
        <button @click="updateUsername">Save</button>
      </div>
    </div>

    <!-- Pulsante per caricare una nuova foto -->
    <input type="file" @change="uploadPhoto" style="display: none;" ref="photoInput" />
    <button @click="triggerFileInput">Upload Photo</button>

    <div v-if="loading">Loading...</div>
    <div v-if="errormsg">{{ errormsg }}</div>

    <div>
      <p>Photos uploaded: {{ photos.length }}</p>
      <p>Followers: <span @click="showFollowers">{{ followers.length }}</span></p>
      <p>Following: <span @click="showFollowing">{{ following.length }}</span></p>
    </div>

    <div v-if="showFollowerList">
      <h2>Followers</h2>
      <ul>
        <li v-for="follower in followers" :key="follower.userId">{{ follower.username }}</li>
      </ul>
    </div>
    
    <div v-if="showFollowingList">
      <h2>Following</h2>
      <ul>
        <li v-for="followee in following" :key="followee.userId">{{ followee.username }}</li>
      </ul>
    </div>

    <div v-for="photo in photos" :key="photo.photoId" class="photo">
      <div class="photo-container">
        <img :src="photo.imageSrc" alt="photo" class="photo-img" @error="handleImageError" />
      </div>
      <div class="photo-info">
        <p>Uploaded at: {{ new Date(photo.uploadDateTime).toLocaleString() }}</p>
        <p>Likes: {{ photo.likes.length }}</p>
        <p>Comments: {{ photo.comments.length }}</p>
      </div>
      <div class="photo-actions">
        <button @click="toggleLike(photo)">Like</button>
        <button @click="commentPhoto(photo.photoId)">Comment</button>
        <button @click="deletePhoto(photo.photoId)">Delete Photo</button>
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
      username: '',
      followers: [],
      newUsername: '', // Nuovo nome utente
      editingUsername: false, // Stato per mostrare/nascondere il campo di modifica nome
      following: [],
      showFollowerList: false,
      showFollowingList: false
    };
  },
  methods: {
    async fetchProfile() {
      this.loading = true;
      try {
        let response = await axios.get(`/user/${this.userId}/profile`);
        this.username = response.data.username; // Assicurati che il campo "username" esista nella risposta
        this.photos = response.data.photos.map(photo => ({
          ...photo,
          comments: photo.Comments || [],
          likes: photo.Likes || []
        }));
        this.followers = response.data.followers;
        this.following = response.data.following;

        // Load images for each photo
        for (let photo of this.photos) {
          this.loadImage(photo.photoId);
        }
      } catch (e) {
        console.error("Error fetching profile:", e);
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
        this.fetchProfile();
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
        this.fetchProfile();
      } catch (e) {
        console.error("Error commenting on photo:", e);
        this.errormsg = e.toString();
      }
    },
    async deletePhoto(photoId) {
      try {
        await axios.delete(`/user/${this.userId}/photos/${photoId}`);
        this.fetchProfile();
      } catch (e) {
        console.error("Error deleting photo:", e);
        this.errormsg = e.toString();
      }
    },
    async deleteComment(photoId, commentId) {
      try {
        await axios.delete(`/user/${this.userId}/photos/${photoId}/comments/${commentId}`);
        this.fetchProfile();
      } catch (e) {
        console.error("Error deleting comment:", e);
        this.errormsg = e.toString();
      }
    },
    async updateUsername() {
      try {
        const response = await axios.put(`/user/${this.userId}/username`, {
          username: this.newUsername
        });
        this.username = response.data.username; // Aggiorna il nome visualizzato
        this.newUsername = ''; // Reset del campo di input
        this.editingUsername = false; // Chiude la modalità di modifica
      } catch (error) {
        console.error("Error updating username:", error);
        this.errormsg = error.toString();
      }
    },
    logout() {
      localStorage.removeItem('token');
      this.$router.replace('/login');
    },
    showFollowers() {
      this.showFollowerList = !this.showFollowerList;
      this.showFollowingList = false;
    },
    showFollowing() {
      this.showFollowingList = !this.showFollowingList;
      this.showFollowerList = false;
    },
    // Metodo per triggerare il file input nascosto
    triggerFileInput() {
      if (this.$refs.photoInput) {
        console.log("Pulsante Upload cliccato, apri il file input");
        this.$refs.photoInput.click(); // Simula il click sul file input nascosto
      } else {
        console.error("Riferimento al file input non trovato");
      }
    },

    // Metodo per gestire l'upload della foto
    async uploadPhoto(event) {
      const file = event.target.files[0];
      if (!file) {
        this.errormsg = 'Nessun file selezionato';
        return;
      }

      // Estrai l'estensione del file
      const allowedExtensions = ['jpg', 'jpeg', 'png', 'gif'];
      const fileExtension = file.name.split('.').pop().toLowerCase();

      // Verifica se l'estensione è tra quelle consentite
      if (!allowedExtensions.includes(fileExtension)) {
        this.errormsg = 'Il file selezionato non è un\'immagine valida. Formati supportati: jpg, jpeg, png, gif';
        console.error(this.errormsg);
        return;
      }

      const formData = new FormData();
      formData.append('file', file);

      try {
        const response = await axios.post(`/user/${this.userId}/photos/`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': `Bearer ${this.userId}`
          }
        });
        console.log('Upload successful:', response.data);
        this.fetchProfile(); // Aggiorna il profilo dopo l'upload
      } catch (error) {
        if (error.response) {
          console.error('Errore API:', error.response.data);
          this.errormsg = error.response.data.message || 'Errore durante l\'upload della foto';
        } else {
          console.error('Errore sconosciuto:', error);
          this.errormsg = 'Errore sconosciuto durante l\'upload della foto';
        }
      }
    }

  },


  mounted() {
    this.fetchProfile();
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