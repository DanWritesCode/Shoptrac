<template>
  <div class="row align-items-center align-content-center py-2" :class="centerText ? 'text-center' : ''">
    <div class="col" v-for="x in data">
      <p class="d-inline">{{x}}</p>
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
    colorClass: {
      type: String,
      default: 'is-red'
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
