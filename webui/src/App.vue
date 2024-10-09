<template>
  <header v-if="showNavBar" class="navbar navbar-expand-lg navbar-dark bg-gradient-blue p-3 shadow">
    <div class="container-fluid">
      <span class="navbar-brand text-white fw-bold">WASAPhoto</span>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <RouterLink to="/" class="nav-link">
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
              Home
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink :to="`/user/${userID}/profile`" class="nav-link">
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
              Profile
            </RouterLink>
          </li>
        </ul>
        <button @click="logout" class="btn btn-outline-light ms-3">Logout</button>
      </div>
    </div>
  </header>

  <div class="container-fluid">
    <div class="row">
      <main class="col-12 px-md-4">
        <RouterView />
      </main>
    </div>
  </div>
</template>


<script>
import { RouterLink, RouterView } from 'vue-router'

export default {
  data() {
    return {
      showNavBar: this.$route.path !== '/login',
      userID: localStorage.getItem('token'),
    }
  },
  watch: {
    $route(to) {
      this.showNavBar = to.path !== '/login';
      this.userID = localStorage.getItem('token');
    },
  },
  methods: {
    logout() {
      localStorage.clear();
      this.$router.replace("/login");
    },
  },
};
</script>

<style scoped>
/* Sfondo a gradiente blu per la navbar */
.bg-gradient-blue {
  background: linear-gradient(135deg, #00b4db, #0083b0);
}

/* Stile per il titolo WASAPhoto */
.navbar-brand {
  font-size: 1.5rem;
  color: white !important;
  cursor: default;
  pointer-events: none;
  background: none !important;
  box-shadow: none !important;
  padding: 0 !important;
  margin: 0 !important;
  border: none !important;
}

/* Stile per i link della navbar */
.nav-link {
  color: white !important;
  font-size: 1.2rem; /* Aumenta la dimensione del testo */
  transition: color 0.3s;
}
.nav-link:hover {
  color: #dcdcdc !important;
}

/* Stile per le icone nei link */
.nav-link i {
  font-size: 1.3rem;
}

/* Stile per il pulsante logout */
.btn-outline-light {
  border-color: white;
  color: white;
  transition: background-color 0.3s, color 0.3s;
}
.btn-outline-light:hover {
  background-color: white;
  color: #0083b0;
}
</style>
