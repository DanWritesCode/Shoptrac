<template>
  <div class="row align-items-center align-content-center py-2" :class="centerText ? 'text-center' : ''">
    <div class="col" v-for="x in data">
      <p class="d-inline fs-16" :class="color">{{x}}</p>
    </div>

  </div>
</template>

<script>
import Dropdown from './Dropdown'

export default {
  name: 'SuInfoItem',
  components: { Dropdown },
  props: {
    name: {
      type: String,
      default: ''
    },
    color: {
      type: String,
      default: 'is-white'
    },
    centerText: {
      type: Boolean,
      default: true
    },
    data: {
      type: Object,
      default: {}
    }
  },
  data() {
    return {
      isToggled: false,
      statusMessage: '',
    }
  },
  methods: {

  }
}
</script>
