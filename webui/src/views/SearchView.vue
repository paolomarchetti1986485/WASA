<template>
  <div>
    <h1>Risultati della ricerca per "{{ username }}"</h1>

    <div v-if="loading">Caricamento...</div>
    <div v-if="errormsg">{{ errormsg }}</div>

    <ul v-if="users.length > 0">
      <li v-for="user in users" :key="user.userId">
        <!-- Reindirizza al profilo utente -->
        <router-link :to="`/user/${user.userId}/nonlogged-profile`">{{ user.username }}</router-link>
      </li>
    </ul>
    <div v-else> Nessun utente trovato.</div>
  </div>
</template>

<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      username: '', // Mantieni il nome utente cercato
      users: [],
      loading: false,
      errormsg: null,
    };
  },
  methods: {
    async searchUsers() {
      this.loading = true;
      this.errormsg = null;
      try {
        // Chiamata API per cercare gli utenti
        const response = await axios.get(`/user/?username=${this.username}`);
        this.users = response.data;
      } catch (error) {
        console.error('Errore durante la ricerca:', error);
        this.errormsg = 'Errore durante la ricerca degli utenti';
      } finally {
        this.loading = false;
      }
    }
  },
  mounted() {
    // Recupera il parametro "username" dalla query string e avvia la ricerca
    this.username = this.$route.query.username || ''; // Default vuoto se non esiste
    if (this.username) {
      this.searchUsers(); // Esegui la ricerca al montaggio della view
    }
  },
  watch: {
    // Rerun the search when the username in the query string changes
    '$route.query.username'(newUsername) {
      this.username = newUsername;
      if (this.username) {
        this.searchUsers();
      }
    }
  }
};
</script>

<style scoped>
/* Aggiungi qui eventuali stili */
</style>
