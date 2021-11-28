<template>
  <button
    :class="{ 'is-active': isToggled }"
    @click="toggleSidebar"
    class="hamburger hamburger--slider cursor"
    type="button"
    aria-label="Menu"
    aria-controls="navigation"
  >
    <span class="hamburger-box">
      <span class="hamburger-inner"></span>
    </span>
  </button>
</template>

<script>
export default {
  name: 'SidebarButton',
  data() {
    return {
      isToggled: false,
      windowWidth: window.innerWidth
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
      window.$nuxt.$emit('toggle-sidebar')
    }
  }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.hamburger {
  display: inline-block;
  cursor: pointer;
  transition-property: opacity, filter;
  transition-duration: 0.15s;
  transition-timing-function: linear;
  font: inherit;
  color: inherit;
  text-transform: none;
  background-color: transparent;
  border: 0;
  margin: 0;
  overflow: visible;
  &:focus {
    outline: none !important;
  }
  &:hover {
    opacity: 0.7;
  }

  &.is-active {
    .hamburger-inner {
      background-color: white;

      &::before,
      &::after {
        background-color: white;
      }
    }
  }
}

.hamburger-box {
  width: 40px;
  height: 24px;
  display: inline-block;
  position: relative;
}

.hamburger-inner {
  display: block;
  top: 50%;
  margin-top: -2px;
  width: 30px;
  height: 2px;
  background-color: $brand;
  border-radius: 4px;
  position: absolute;
  transition-property: transform;
  transition-duration: 0.15s;
  transition-timing-function: ease;

  &::before,
  &::after {
    width: 30px;
    height: 2px;
    background-color: $brand;
    border-radius: 4px;
    position: absolute;
    transition-property: transform;
    transition-duration: 0.15s;
    transition-timing-function: ease;
  }

  &::before,
  &::after {
    content: '';
    display: block;
  }

  &::before {
    top: -10px;
  }

  &::after {
    bottom: -10px;
  }
}

.hamburger--slider {
  .hamburger-inner {
    top: 2px;

    &::before {
      top: 10px;
      transition-property: transform, opacity;
      transition-timing-function: ease;
      transition-duration: 0.15s;
    }

    &::after {
      top: 20px;
    }
  }

  &.is-active .hamburger-inner {
    transform: translate3d(0, 10px, 0) rotate(45deg);

    &::before {
      transform: rotate(-45deg) translate3d(-5.71429px, -6px, 0);
      opacity: 0;
    }

    &::after {
      transform: translate3d(0, -20px, 0) rotate(-90deg);
    }
  }
}
</style>
