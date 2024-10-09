<template>
  <div>
    <div v-if="isBannedByUser">
      <p>Non hai accesso a questo profilo perché sei stato bannato dall'utente.</p>
    </div>
    <div v-else>
      <h1 v-if="username">{{ username }}'s Profile Page</h1>
      <h1 v-else>Loading...</h1>
      <div v-if="!isOwner">
        <button @click="toggleFollow">
          {{ isFollowing ? 'Unfollow' : 'Follow' }}
        </button>
        <!-- Pulsante Ban/Unban -->
        <button @click="toggleBan">
          {{ hasBanned ? 'Unban' : 'Ban' }}
        </button>

      </div>
      <button @click="logout">Logout</button>
      
      <div v-if="loading">Loading...</div>
      <div v-if="errormsg">{{ errormsg }}</div>
      <div v-if="photos.length === 0 && !loading" class="empty-profile-message">
        Nessun post disponibile
      </div>

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
        </div>
        <div v-for="comment in photo.comments" :key="comment.commentId" class="comment">
          <p>{{ comment.commentText }}</p>
          <button v-if="loggedInUserId && comment.userId == loggedInUserId" @click="deleteComment(photo.photoId, comment.commentId)">
            Delete
          </button>
        </div>
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
      userId: this.$route.params.userId || localStorage.getItem('token'),
      username: '',
      followers: [],
      following: [],
      showFollowerList: false,
      showFollowingList: false,
      isFollowing: false,
      isBannedByUser: false,
      hasBanned: false  // Se l'utente loggato ha bannato l'utente visualizzato


    };
  },
  methods: {
    async fetchProfile() {
    this.loading = true;
    try {
        // Verifica se l'utente loggato è stato bannato dall'utente visualizzato
        let banCheckResponse = await axios.get(`/user/${this.userId}/ban/${this.loggedInUserId}`);
        if (banCheckResponse.data.isBanned) {
            this.isBannedByUser = true;
            this.loading = false;
            return; // Interrompi il caricamento del profilo
        }

        let response = await axios.get(`/user/${this.userId}/profile`);
        if (response.data) {
          this.username = response.data.username; // Assicurati che il campo "username" esista nella risposta
          this.photos = Array.isArray(response.data.photos) ? response.data.photos.map(photo => ({
            ...photo,
            comments: photo.Comments || [],
            likes: photo.Likes || []
          })) : [];
          this.followers = response.data.followers || [];
          this.following = response.data.following || [];
          this.isFollowing = this.followers.some(follower => follower.userId == this.loggedInUserId);

          // Verifica se l'utente loggato ha bannato l'utente visualizzato
          let hasBannedResponse = await axios.get(`/user/${this.loggedInUserId}/ban/${this.userId}`);
          this.hasBanned = hasBannedResponse.data.isBanned;

          // Carica le immagini per ogni foto
          for (let photo of this.photos) {
            this.loadImage(photo.photoId);
          }
        } else {
          this.photos = [];
          this.followers = [];
          this.following = [];
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
      const loggedInUserId = localStorage.getItem('token');
      console.log("Logged in User ID:", loggedInUserId); // Log del token

      const hasLiked = photo.likes.some(like => like.userId == loggedInUserId);
      try {
        if (hasLiked) {
          await axios.delete(`/photos/${photo.photoId}/likes/${loggedInUserId}`);
        } else {
          await axios.put(`/photos/${photo.photoId}/likes/${loggedInUserId}`);
        }
        this.fetchProfile();
      } catch (e) {
        console.error("Error toggling like on photo:", e); // Log dell'errore
        this.errormsg = e.toString();
      }
    },

    async commentPhoto(photoId) {
      const loggedInUserId = localStorage.getItem('token'); // Usa l'ID dell'utente loggato dal token
      const text = prompt("Enter your comment:");
      if (!text) return;
      try {
        await axios.post(`/user/${loggedInUserId}/photos/${photoId}/comments/`, { text });
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
      console.log("Logged in user ID:", this.loggedInUserId); // Log dell'utente loggato
      try {
        await axios.delete(`/user/${this.loggedInUserId}/photos/${photoId}/comments/${commentId}`);
        this.fetchProfile();
      } catch (e) {
        console.error("Error deleting comment:", e.response || e.message);
        this.errormsg = e.toString();
      }
    },
    async toggleBan() {
      try {
        if (this.hasBanned) {
          // Se l'utente è già bannato, sbannalo
          await axios.delete(`/user/${this.loggedInUserId}/ban/${this.userId}`);
          this.hasBanned = false;
        } else {
          // Se l'utente non è bannato, bannalo
          await axios.put(`/user/${this.loggedInUserId}/ban/${this.userId}`);
          this.hasBanned = true;
        }
      } catch (e) {
        console.error("Error toggling ban:", e);
        this.errormsg = e.toString();
      }
    },
    async toggleFollow() {
      try {
        if (this.isFollowing) {
          // Unfollow l'utente
          await axios.delete(`/user/${this.loggedInUserId}/follow/${this.userId}`);
          this.isFollowing = false;
        } else {
          // Follow l'utente
          await axios.put(`/user/${this.loggedInUserId}/follow/${this.userId}`);
          this.isFollowing = true;
        }
      } catch (e) {
        console.error("Error toggling follow:", e);
        this.errormsg = e.toString();
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
    }
  },
  mounted() {
    this.loggedInUserId = localStorage ? localStorage.getItem('token') : null;
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