<template>
  <div class="row align-items-center align-content-center py-2">
    <div class="col-5">
      <p :class="colorClass + '-text'" class="d-inline font-weight-bold fs-14">
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
    colorClass: {
      type: String,
      default: 'is-red'
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
