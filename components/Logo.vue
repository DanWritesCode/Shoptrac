<template>
  <svg
    :fill="color"
    :class="hoverClass"
    width="50"
    height="50"
    viewBox="0 0 50 50"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path d="M29.3851 0H18.4204V3.56715H29.3851V0Z"></path>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M49.5716 15.0333C48.9948 12.7619 47.9102 10.8954 46.4812 9.41309L44.3787 10.5693C45.6714 11.9629 46.7098 13.7971 47.4691 16.1894L49.5716 15.0333Z"
    ></path>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M47.2134 14.1516L45.9623 11.8767L42.8367 13.5954L44.0876 15.8704L47.2134 14.1516Z"
    ></path>
    <path d="M25.6863 3.77808H22.1191V47.5938H25.6863V3.77808Z"></path>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M45.8101 27.4691V23.9021H1.99438V27.4691H45.8101Z"
    ></path>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M23.9025 1.78369C37.0946 1.78369 47.8049 12.494 47.8049 25.6862C47.8049 38.8783 37.0946 49.5886 23.9025 49.5886C10.7103 49.5886 0 38.8783 0 25.6862C0 12.494 10.7103 1.78369 23.9025 1.78369ZM23.9025 5.60813C34.9839 5.60813 43.9805 14.6048 43.9805 25.6862C43.9805 36.7676 34.9839 45.7642 23.9025 45.7642C12.8211 45.7642 3.82444 36.7676 3.82444 25.6862C3.82444 14.6048 12.8211 5.60813 23.9025 5.60813Z"
    ></path>
    <path
      fill-rule="evenodd"
      clip-rule="evenodd"
      d="M23.9023 12.8761C30.9723 12.8761 36.7122 18.616 36.7122 25.6861C36.7122 32.7561 30.9723 38.4961 23.9023 38.4961C16.8322 38.4961 11.0923 32.7561 11.0923 25.6861C11.0923 18.616 16.8322 12.8761 23.9023 12.8761ZM23.9023 16.3348C29.0634 16.3348 33.2536 20.525 33.2536 25.6861C33.2536 30.8472 29.0634 35.0374 23.9023 35.0374C18.7412 35.0374 14.551 30.8472 14.551 25.6861C14.551 20.525 18.7412 16.3348 23.9023 16.3348Z"
    ></path>
  </svg>
</template>

<script>
export default {
  name: 'Logo',
  props: {
    color: {
      type: String,
      default: '#4A4754'
    },
    colorHover: {
      type: String,
      default: '#EF423C'
    },
    hover: {
      type: Boolean
    }
  },
  data() {
    return {
      currentColor: '#4A4754'
    }
  },
  computed: {
    hoverClass() {
      if (this.hover) {
        return 'fillable'
      }
      return ''
    }
  }
}
</script>

<style lang="scss">
.fillable {
  transition: all 150ms ease-in;
  &:hover {
    transition: all 150ms ease-out;
    fill: #ef423c;
  }
}
</style>
