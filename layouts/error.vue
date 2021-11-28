<template>
  <div class="container">
    <h1>An error occurred (Error {{ error.statusCode }})</h1>
  </div>
</template>

<script>
export default {
  props: ['error'],
  layout: 'default'
}
</script>