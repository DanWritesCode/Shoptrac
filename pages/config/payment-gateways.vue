<template>
  <div class="pb-5">
    <div class="row">
      <!-- Double mt-4 because the above boxes have weird sizing issues. Oh well... -->
      <div class="col-12 mt-4"><h3>Data...</h3></div>
    </div>
    <div class="row">

    </div>
  </div>
</template>

<script>

export default {
  name: 'Payment Gateways',
  components: { },
  layout: 'dashboard',
  transition: 'fade',
  head() {
    return {
      title: 'Shoptrac | Config - Payment Gateways'
    }
  },
  data() {
    return {
      notices: [],
    }
  },
  methods: {

  }
}
</script>