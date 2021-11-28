<template>
  <div class="sidebar-footer ml-4 mt-n2">

  </div>
</template>

<script>
export default {
  name: 'SidebarFooter',
  methods: {
    logout() {
      localStorage.setItem('loggedIn', false)
      localStorage.setItem('username', '')
      localStorage.setItem('token', '')
      this.$router.push('/')
    }
  }
}
</script>

<style lang="scss">
.sidebar {
  .sidebar-footer {
  }
  .hover-right {
    svg {
      transition: all 250ms ease-in;
    }
    &:hover {
      svg {
        transition: all 250ms ease-out;
        margin-left: 95px;
      }
    }
  }
}
</style>
