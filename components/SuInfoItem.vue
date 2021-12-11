<template>
  <div class="row align-items-center align-content-center py-2">
    <div class="col-5">
      <p :class="currentColor + '-text'" class="d-inline font-weight-bold fs-14">
        {{ name }}
      </p>
    </div>
    <div class="col-4">
      <p class="d-inline">Test</p>
    </div>
    <div class="col-3">
      <p class="d-inline">Test</p>
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
    time: {
      type: Number,
      default: 0
    },
    status: {
      type: String,
      default: 'is-green'
    }
  },
  data() {
    return {
      isToggled: false,
      currentColor: 'is-red',
      statusMessage: '',
    }
  },
  methods: {

  }
}
</script>
