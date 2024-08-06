<template>
  <div>
    <h1 v-if="username">{{ username }}'s Profile Page</h1>
    <h1 v-else>Loading...</h1>
    <button @click="logout">Logout</button>
    
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
