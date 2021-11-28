<template>
  <div class="menu">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'Dropdown'
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.menu {
  position: absolute;
  background: #08070a;
  box-shadow: 2px 2px 7px 0px rgba(5, 5, 5, 0.5);
  left: 0;
  right: 0;
  top: 0;
  margin-top: 25px;
  z-index: 1000;
  border-radius: 3px;
  text-align: center;
  width: 155px;
  padding: 3px;
  ul {
    margin: 0;
    li {
      padding: 7px 10px;
      &:not(:last-of-type) {
        border-bottom: 1px solid $dark;
      }
      a {
        color: $text;
        &:hover {
          color: $brand;
        }
      }
    }
  }
}
</style>
