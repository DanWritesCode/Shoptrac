<template>
  <div id="app">
    <slide-x-left-transition orign="left">
      <db-sidebar></db-sidebar>
    </slide-x-left-transition>
    <div id="content">
      <db-navbar></db-navbar>
      <div id="wrapper">
        <div class="container-fluid p-4">
          <nuxt />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SlideXLeftTransition from 'vue2-transitions/src/Slide/SlideXLeftTransition'
import DbSidebar from '../components/Sidebar'
import DbNavbar from '../components/HorizontalNavbar'
export default {
  name: 'Dashboard',
  components: { SlideXLeftTransition, DbNavbar, DbSidebar },
  middleware: 'authenticated'
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css?family=Poppins:200,300,400,600,700,800');
@import '../assets/sass/variables';

body {
  font-family: 'Poppins', sans-serif;
  color: $text;
  font-size: 14px;
  background: $bg;
}
#app {
  display: flex;
  align-items: stretch;
  position: relative;
}

#content {
  min-height: 100vh;
  width: 100vw;
  transition: all 0.3s;
}

#wrapper {
  padding: 10px;
}
</style>
