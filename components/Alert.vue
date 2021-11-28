<template>
  <div :class="type" v-if="isOpen" class="alert">
    <div class="row align-items-center">
      <div class="col-10">
        <slot></slot>
      </div>
      <div class="col-2 text-right">
        <span @click="Close()" class="text-white cursor">&times;</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Alert',
  props: {
    type: {
      type: String,
      default: 'is-green'
    }
  },
  data() {
    return {
      isOpen: true
    }
  },
  methods: {
    Close() {
      this.isOpen = !this.isOpen
    }
  }
}
</script>
