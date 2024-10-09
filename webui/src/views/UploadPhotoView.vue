<template>
  <div class="upload-photo">
    <h2>Upload a New Photo</h2>
    <input type="file" @change="handleFileUpload" />
    <button @click="uploadPhoto">Upload</button>
    <div v-if="errormsg" class="error">{{ errormsg }}</div>
    <button @click="goBack">Cancel</button>
  </div>
</template>

<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      selectedFile: null,
      errormsg: null,
      userId: localStorage.getItem('token') // Prendi userId dal token salvato
    };
  },
  methods: {
    handleFileUpload(event) {
      this.selectedFile = event.target.files[0];
      this.errormsg = null;
    },
    // Metodo per gestire l'upload della foto
    async uploadPhoto() {
      if (!this.selectedFile) {
        this.errormsg = 'Nessun file selezionato';
        return;
      }

      // Estrai l'estensione del file
      const allowedExtensions = ['jpg', 'jpeg', 'png', 'gif'];
      const fileExtension = this.selectedFile.name.split('.').pop().toLowerCase();

      // Verifica se l'estensione è tra quelle consentite
      if (!allowedExtensions.includes(fileExtension)) {
        this.errormsg = 'Il file selezionato non è un\'immagine valida. Formati supportati: jpg, jpeg, png, gif';
        console.error(this.errormsg);
        return;
      }

      const formData = new FormData();
      formData.append('file', this.selectedFile);

      try {
        const response = await axios.post(`/user/${this.userId}/photos/`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': `Bearer ${this.userId}`
          }
        });
        console.log('Upload successful:', response.data);
        this.$router.push(`user/${this.userId}/profile`); // Reindirizza al profilo dopo l'upload
      } catch (error) {
        if (error.response) {
          console.error('Errore API:', error.response.data);
          this.errormsg = error.response.data.message || 'Errore durante l\'upload della foto';
        } else {
          console.error('Errore sconosciuto:', error);
          this.errormsg = 'Errore sconosciuto durante l\'upload della foto';
        }
      }
    },
    goBack() {
      this.$router.push(`user/${this.userId}/profile`);
    }
  }
};
</script>

<style>
.upload-photo {
  padding: 20px;
}
.error {
  color: red;
  margin-top: 10px;
}
</style>
