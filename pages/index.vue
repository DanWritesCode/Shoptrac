<template>
  <div>
    <div class="py-5">
      <home-header :users="users" :successful="successful"></home-header>
    </div>
    <div class="pt-5">
    </div>
  </div>
</template>

<script>
import HomeHeader from '../components/HomeHeader'

export default {
  name: 'Home',
  components: { HomeHeader },
  transition: 'fade',
  head() {
    return {
      title: 'Shoptrac | Home'
    }
  },
  data() {

  },
  async mounted() {
    if (
        localStorage.getItem('username') === null ||
          localStorage.getItem('token') === null || localStorage.getItem('loggedIn') === null) {

      localStorage.setItem('username', '')
      localStorage.setItem('token', '')
      localStorage.setItem('loggedIn', false)
    }
  }

}
</script>
