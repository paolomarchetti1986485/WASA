<template>
  <div class="container-fluid h-100 m-0 p-0 login">
    <div class="row">
      <div class="col">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
      </div>
    </div>

    <div class="row h-100 w-100 m-0">
      <form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0">
        <div class="row mt-2 mb-3 border-bottom">
          <div class="col">
            <h2 class="login-title">Login</h2>
          </div>
        </div>

        <div class="row mt-2 mb-3">
          <div class="col">
            <input 
              type="text" 
              class="form-control" 
              v-model="identifier" 
              maxlength="16"
              minlength="3"
              placeholder="Enter your username" 
            />
          </div>
        </div>

        <div class="row mt-2 mb-5">
          <div class="col">
            <button class="btn btn-dark" :disabled="isButtonDisabled"> 
              Login 
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from '../services/axios'; // Assicurati che il percorso sia corretto

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

<style>


.login-title {
  color: black;
}
</style>
