<template>
  <div v-if="isToggled" class="sidebar">
    <db-sidebar-header></db-sidebar-header>
    <div class="sidebar-wrapper">
      <db-sidebar-body></db-sidebar-body>
      <db-sidebar-footer></db-sidebar-footer>
    </div>
  </div>
</template>

<script>
import DbSidebarHeader from './SidebarHeader'
import DbSidebarBody from './SidebarBody'
import DbSidebarFooter from './SidebarFooter'
export default {
  name: 'Sidebar',
  components: { DbSidebarFooter, DbSidebarBody, DbSidebarHeader },
  data() {
    return {
      windowWidth: window.innerWidth,
      isToggled: false
    }
  },
  watch: {
    windowWidth() {
      if (this.windowWidth > 990) {
        this.isToggled = true
      }
    }
  },
  mounted() {
    window.$nuxt.$on('toggle-sidebar', () => {
      this.toggleSidebar()
    })
    if (this.windowWidth > 990) {
      this.isToggled = true
    }
    this.$nextTick(() => {
      window.addEventListener('resize', () => {
        this.windowWidth = window.innerWidth
      })
    })
  },
  methods: {
    toggleSidebar() {
      this.isToggled = !this.isToggled
    }
  }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.sidebar {
  width: 300px;
  background-color: $dark;
  .sidebar-wrapper {
    position: sticky;
    top: 10px;
  }
  -webkit-box-shadow: 0 2px 22px 0 rgba(0,0,0,.2),0 2px 30px 0 rgba(0,0,0,.35);
  box-shadow: 0 2px 22px 0 rgba(0,0,0,.2),0 2px 30px 0 rgba(0,0,0,.35);
}
</style>
