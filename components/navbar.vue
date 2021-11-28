<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-transparent">
    <div class="container">
      <div class="navbar-header d-inline-block">
        <n-link class="navbar-brand" to="/">
          <div class="d-inline-block" style="height: 100%; vertical-align: center;"><img src="~/assets/images/logo.png" /></div>
          <div class="d-inline-block" style="height: 100%; vertical-align: center;">
            <p class="mb-0">StonksUp</p>
          </div>
        </n-link>
      </div>
      <button
        :class="{ 'is-active': isToggled }"
        @click="toggleNavbar()"
        class="hamburger hamburger-vortex d-block d-lg-none"
        type="button"
      >
        <span class="hamburger-box">
          <span class="hamburger-inner"></span>
        </span>
      </button>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav mr-auto"></ul>
        <ul class="navbar-nav ml-auto align-items-center">
          <li class="nav-item">
            <n-link class="nav-link" to="/">Home</n-link>
          </li>
          <li class="nav-item">
            <n-link class="nav-link" to="/about">About</n-link>
          </li>
          <li class="nav-item">
            <n-link class="nav-link" to="/faq">FAQ</n-link>
          </li>
          <li v-if="!isLoggedin" class="nav-item">
            <n-link class="nav-link" to="/sign-in">Sign in</n-link>
          </li>
          <li v-if="!isLoggedin" class="nav-item">
            <n-link class="nav-link btn btn-red" to="/sign-up">Sign up</n-link>
          </li>
          <li v-if="isLoggedin" class="nav-item">
            <n-link class="nav-link btn btn-red" to="/dashboard">
              Dashboard
            </n-link>
          </li>
        </ul>
      </div>
      <transition name="fade">
        <div v-if="isToggled" class="d-block d-lg-none nav-mobile">
          <ul class="navbar-nav">
            <li class="nav-item">
              <n-link class="nav-link" to="/">Home</n-link>
            </li>
            <li class="nav-item">
              <n-link class="nav-link" to="/about">About</n-link>
            </li>
            <li v-if="!isLoggedin" class="nav-item">
              <n-link class="nav-link" to="/sign-in">Sign in</n-link>
            </li>
            <li v-if="!isLoggedin" class="nav-item">
              <n-link class="nav-link btn btn-red" to="/sign-up"
                >Sign up</n-link
              >
            </li>
            <li v-if="isLoggedin" class="nav-item">
              <n-link class="nav-link btn btn-red" to="/dashboard">
                Dashboard
              </n-link>
            </li>
          </ul>
        </div>
      </transition>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      isToggled: false,
      isLoggedin: false
    }
  },
  mounted() {
    if (localStorage.getItem('loggedIn') === 'true') {
      this.isLoggedin = true
    }
  },
  methods: {
    toggleNavbar() {
      this.isToggled = !this.isToggled
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/sass/variables';
@import '../assets/sass/components/core/navbar';
</style>
