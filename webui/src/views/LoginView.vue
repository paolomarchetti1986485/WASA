<template>
  <div class="container-fluid h-100 m-0 p-0 login">
    <div class="row h-100 m-0 d-flex align-items-center justify-content-center">
      <form @submit.prevent="login" class="login-form d-flex flex-column align-items-center p-4 rounded shadow">
        <h1 class="app-title mb-4">WASAPhoto</h1>
        <ErrorMsg v-if="errormsg" :msg="errormsg" class="mb-3"></ErrorMsg>

        <input 
          type="text" 
          class="form-control mb-3 login-input" 
          v-model="identifier" 
          maxlength="16"
          minlength="3"
          placeholder="Enter your username" 
        />

        <button class="btn login-btn w-100" :disabled="isButtonDisabled"> 
          Login 
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      errormsg: null,
      identifier: "",
    };
  },
  computed: {
    isButtonDisabled() {
      return this.identifier == null || this.identifier.length > 16 || this.identifier.length < 3 || this.identifier.trim().length < 3;
    },
  },
  methods: {
    async login() {
      this.errormsg = null;
      try {
        console.log("Sending request to server with identifier:", this.identifier);
        let response = await axios.post("/session", {
          username: this.identifier.trim(),
        });

        console.log("Received response from server:", response.data);

        localStorage.setItem('token', response.data.userId);
        this.$router.replace("/");
        this.$emit('updatedLoggedChild', true);
      } catch (e) {
        console.error("Error during login request:", e);
        this.errormsg = e.toString();
      }
    },
  },
  mounted() {
    if (localStorage.getItem('token')) {
      this.$router.replace("/");
    }
  },
};
</script>

<style scoped>
/* Sfondo a gradiente per tutta la pagina */
.login {
  background: linear-gradient(135deg, #00b4db, #0083b0);
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Box di login */
.login-form {
  background: white;
  max-width: 400px;
  width: 100%;
  border-radius: 8px;
  padding: 30px;
}

/* Titolo dell'app */
.app-title {
  font-size: 2rem;
  color: #0083b0;
  text-align: center;
  font-weight: bold;
}

/* Stile dell'input */
.login-input {
  border: 1px solid #ddd;
  padding: 12px;
  border-radius: 4px;
  font-size: 1rem;
}

/* Pulsante di login */
.login-btn {
  background-color: #0083b0;
  color: white;
  border: none;
  padding: 12px;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-btn:hover {
  background-color: #006994;
}

.login-btn:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

/* Messaggio di errore */
.error-msg {
  color: red;
  font-size: 0.9rem;
}
</style>
