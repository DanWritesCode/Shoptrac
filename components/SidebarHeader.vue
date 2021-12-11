<template>
  <div class="sidebar-header">
    <h1>
      <n-link to="/">
        <img
          src="~/assets/images/logo.png"
          alt="StonksUp"
          class="logo"
        />
        StonksUp
      </n-link>
    </h1>
    <p class="text-grey font-weight-bold fs-14">
      <a href="" class="text-blue"> Welcome, Username </a>
    </p>
  </div>
</template>c

<script>
export default {
  name: 'SidebarHeader',
  data() {
    return {
      username: localStorage.getItem('username')
    }
  },
  mounted() {

  },
  watched: {

  }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.sidebar {
  .sidebar-header {
    text-align: center;
    display: flex;
    margin-top: 40px;
    flex-direction: column;
    align-items: center;
    border-bottom: $break 1px solid;

    .logo {
      width: 38px;
      object-fit: contain;
      margin-right: 5px;
    }

    .avatar {
      width: 100px;
      border-radius: 10%;
      object-fit: contain;
    }
    h1 {
      font-size: 18px;
      font-family: 'Poppins', sans-serif;
      font-weight: 700;
      margin-bottom: 30px;
      a {
        color: white;
      }
    }
  }
}
</style>
