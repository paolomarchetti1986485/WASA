<template>
  <div>
    <div class="home-container">
      <div class="header-container">
        <h1 class="home-title">Home Page</h1>
        <div class="search-bar mb-4">
          <input
            type="text"
            v-model="searchQuery"
            placeholder="Cerca un utente..."
            @keydown.enter="searchUser"
            class="search-input"
          />
          <button @click="searchUser" class="search-button">Cerca</button>
        </div>
      </div>
    </div>
    <div v-if="loading">Loading...</div>
    <div v-if="errormsg">{{ errormsg }}</div>

    <div v-for="photo in photos" :key="photo.photoId" class="photo-card">
      <div class="photo-header">
        <p @click="goToProfile(photo.userId)" class="user-profile-link">
          {{ photo.username || 'Unknown User' }}
        </p>
        <hr class="user-separator" /> <!-- Barra sotto il nome dell'utente -->
      </div>
      <div class="photo-container">
        <img :src="photo.imageSrc" alt="photo" class="photo-img" @error="handleImageError" />
      </div>
      <div class="photo-details">
        <button @click="toggleLike(photo)" class="like-button">
          ❤︎ {{ photo.likes.length }}
        </button>
        <div class="comments">
          <div v-for="comment in photo.comments" :key="comment.commentId" class="comment-item">
            <p>{{ comment.commentText }}</p>
            <button v-if="comment.userId == userId" @click="deleteComment(photo.photoId, comment.commentId)" class="delete-button">
              Delete
            </button>
          </div>
        </div>
        <button @click="commentPhoto(photo.photoId)" class="comment-button">Comment</button>
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
/* Contenitore principale per la home */
.home-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 20px;
}

/* Stile per il titolo della home */
.home-title {
  font-size: 2rem;
  margin-bottom: 20px;
  text-align: center;
  color: #333;
}
/* Stile per il pulsante della barra di ricerca */
.search-button {
  background-color: #0083b0;
  color: white;
  border: none;
  padding: 10px 15px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.search-button:hover {
  background-color: #005f7a;
}
/* Stile per il contenitore della barra di ricerca */
.search-bar {
  display: flex;
  align-items: center;
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  width: 100%; /* Imposta la larghezza al 100% */
  max-width: 500px; /* Larghezza massima per evitare che sia troppo grande */
  margin: 0 auto; /* Centra la barra di ricerca */
}

/* Stile per l'input della barra di ricerca */
.search-input {
  flex: 1;
  border: none;
  padding: 10px;
  font-size: 1rem;
  border-radius: 0;
}

/* Stile per il separatore sotto il nome dell'utente */
.user-separator {
  border: none;
  border-top: 2px solid #333; /* Colore più scuro per visibilità */
  margin: 8px 0;
}

/* Stile per l'immagine e il contenitore della foto */
.photo-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: auto; /* Altezza automatica per adattarsi all'immagine */
  border: 2px solid black; /* Cornice nera uniforme */
  border-radius: 8px;
  margin: 0 auto 10px auto; /* Centra il contenitore orizzontalmente */
  overflow: hidden;
  max-width: 400px; /* Larghezza massima per centrare la foto */
}

/* Rimuovere l'overflow e assicurare il riempimento dell'immagine */
.photo-img {
  width: 100%;
  height: auto;
  object-fit: contain; /* Mantiene le proporzioni e centra l'immagine */
}

/* Stile per il contenitore delle foto */
.photo-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  margin: 20px auto;
  padding: 10px;
  max-width: 500px;
  background-color: #f9f9f9;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
}

/* Stile per il nome dell'utente */
.photo-header {
  font-weight: bold;
  margin-bottom: 10px;
  cursor: pointer;
}

/* Stile per la sezione dei dettagli della foto */
.photo-details {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

/* Stile per il pulsante di like */
.like-button {
  background-color: transparent;
  border: none;
  cursor: pointer;
  font-size: 1rem;
  margin-bottom: 10px;
  color: #ff5e5e;
}

/* Stile per i commenti */
.comments {
  font-size: 0.9rem;
  color: #555;
  margin-bottom: 10px;
}

/* Stile per il pulsante commento */
.comment-button {
  background-color: #0083b0;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 5px 10px;
  cursor: pointer;
  align-self: flex-start;
}

</style>