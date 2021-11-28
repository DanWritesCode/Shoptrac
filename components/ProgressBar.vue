<template>
  <div>
    <svg
      class="progress-circle"
      viewBox="0 0 106 106"
      version="1.1"
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
    >
      <defs>
        <linearGradient id="grad1" x1="0%" y1="0%" x2="100%" y2="0%">
          <stop offset="0%" style="stop-color:#EF423C;stop-opacity:1"></stop>
          <stop offset="100%" style="stop-color:#FF001F;stop-opacity:1"></stop>
        </linearGradient>
      </defs>
      <g
        id="Page-1"
        stroke="none"
        stroke-width="1"
        fill="none"
        fill-rule="evenodd"
      >
        <text
          x="-67"
          y="57"
          fill="white"
          style="transform: scaleX(-1);font-size: 15px;font-family: 'Inter', sans-serif;font-weight: bold;letter-spacing: 1px;"
        >
          {{ percent }}%
        </text>
        <g id="ProgressBar" transform="translate(-17.000000, -17.000000)">
          <circle
            id="Oval"
            stroke="url(#grad1)"
            stroke-width="5"
            fill-rule="nonzero"
            cx="70"
            cy="70"
            r="50"
          ></circle>
          <path
            id="Oval-Copy"
            :stroke-dasharray="circle"
            class="progress-circle__path"
            d="M70,120 C97.6142375,120 120,97.6142375 120,70 C120,42.3857625 97.6142375,20 70,20 C42.3857625,20 20,42.3857625 20,70 C20,97.6142375 42.3857625,120 70,120 Z"
            stroke="#100F15"
            stroke-width="7.5"
            fill-rule="nonzero"
            transform="translate(70.000000, 70.000000) rotate(-125.000000) translate(-70.000000, -70.000000) "
          ></path>
        </g>
      </g>
    </svg>
  </div>
</template>

<script>
export default {
  name: 'ProgressBar',
  props: {
    percent: {
      type: Number,
      default: 0
    }
  },
  computed: {
    circle() {
      return (this.percent / 100) * 100 * Math.PI + ',9999'
    }
  }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.progress-circle {
  max-width: 70px;
  max-height: 70px;
  width: 100%;
  transform: scaleX(-1) rotate(0deg);

  &__percent {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }

  &__container {
    display: inline-block;
    position: relative;
  }

  &__path {
    transition: 0.5s ease-in-out all;
  }

  &__button {
    text-decoration: none;
    margin-top: 24px;
    margin-left: 24px;
    background: $dark;
    color: white;
    padding: 18px 24px;
    border-radius: 4px;
    transition: 0.25s linear background;

    &:hover {
      cursor: pointer;
      background: darken($dark, 5%);
      transition: 0.25s linear background;
    }
  }
}
</style>
