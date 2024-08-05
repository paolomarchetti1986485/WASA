<template>
  <div>
    <h1>Home Page</h1>
    <button @click="logout">Logout</button>
  </div>
</template>

<script>
export default {
  methods: {
    logout() {
      localStorage.removeItem('token');
      this.$router.replace('/login');
    }
  }
}
</script>