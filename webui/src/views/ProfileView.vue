<template>
  <div>
    <h1 v-if="username">{{ username }}'s Profile Page</h1>
    <h1 v-else>Loading...</h1>
    
    <div v-if="photos.length === 0 && !loading" class="empty-profile-message">
      Non hai aggiunto nessuna foto, comincia a postare!
    </div>

    <div>
      <button @click="editingUsername = !editingUsername">
        {{ editingUsername ? 'Cancel' : 'Update Username' }}
      </button>
      <div v-if="editingUsername">
        <input v-model="newUsername" placeholder="Enter new username" />
        
        <button @click="updateUsername">Save</button>
        <div v-if="errormsg" class="error-message">{{ errormsg }}</div>
      </div>
    </div>

    <!-- Pulsante per navigare alla view di upload delle foto -->
    <button @click="goToUploadPhoto">Upload Photo</button>

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
        if (response.data) {
          this.username = response.data.username; // Assicurati che il campo "username" esista nella risposta

          // Garantisce che photos sia sempre un array
          this.photos = Array.isArray(response.data.photos) ? response.data.photos.map(photo => ({
            ...photo,
            comments: photo.Comments || [],
            likes: photo.Likes || []
          })) : [];

          this.followers = response.data.followers || [];
          this.following = response.data.following || [];

          // Load images for each photo
          if (this.photos.length > 0) {
            for (let photo of this.photos) {
              this.loadImage(photo.photoId);
            }
          }
        } else {
          this.photos = [];
          this.followers = [];
          this.following = [];
        }
      } catch (e) {
        console.error("Error fetching profile:", e);
        this.errormsg = e.response ? e.response.data.message : e.message;
      }
      this.loading = false;
    },
    async checkUsernameAvailability() {
      if (!this.newUsername.trim()) {
        this.usernameExists = false;
        return;
      }

      try {
        const response = await axios.get(`/user/check-username?username=${this.newUsername}`);
        this.usernameExists = response.data.exists; // Se il nome utente esiste già, imposta a true
      } catch (error) {
        // Non loggare l'errore se si tratta di un errore previsto (esempio: il server risponde con 404 se il nome utente è disponibile)
        if (error.response && error.response.status !== 404) {
          console.warn("Errore durante la verifica del nome utente:", error);
        }
      }
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
        this.errormsg = null; // Resetta il messaggio di errore
      } catch (error) {
        if (error.response) {
          if (error.response.status === 409) {
            // Codice HTTP 409: Nome utente già in uso
            this.errormsg = 'Nome utente già in uso.';
          } else {
            // Qualsiasi altro errore
            this.errormsg = 'Errore durante l\'aggiornamento del nome utente.';
          }
        } else {
          // Errore generico se non ci sono risposte dal server
          this.errormsg = 'Errore durante l\'aggiornamento del nome utente.';
        }
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
    goToUploadPhoto() {
    this.$router.push('/upload');
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