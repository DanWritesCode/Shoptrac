<template>
  <div class="modal">
    <div @click="close()" class="modal-background"></div>
    <div class="modal-body">
      <div class="card bg-dark">
        <div
          v-if="hasTitle"
          class="mb-3 px-4 py-3 border-bottom border-grey text-left"
        >
          <h2 class="ff-poppins-sb fs-18 text-grey-alt m-0">
            {{ modalTitle }}
          </h2>
        </div>
        <div class="card-body px-4 pb-4">
          <slot></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SuModal',
  props: {
    modalId: {
      type: String,
      default: ''
    },
    hasTitle: {
      type: Boolean,
      default: true
    },
    modalTitle: {
      type: String,
      default: ''
    }
  },
  methods: {
    close() {
      window.$nuxt.$emit('closeModal', {
        modalId: this.modalId
      })
    }
  }
}
</script>

<style lang="scss">
.modal {
  position: fixed;
  z-index: 999;
  width: 100%;
  height: 100%;
  left: 0;
  right: 0;
  bottom: 0;
  outline: 0;
 overflow: scroll;
  /* padding-top:50%; /* center the top of child elements vetically
  padding-bottom:50%;*/

  @media (min-width: 980px) {
    top: 0;
    justify-content: center;
    align-items: center;
    display: flex;
  }
  .modal-body {
    position: relative;
    min-width: 310px;
    z-index: 1000;
    @media (max-width: 980px) {
      width: 100%;
      display: block;
      vertical-align: middle;
      position: absolute;
      /*transform: translateY(-50%);
      padding-top:100%;*/

    }
    .card {
      box-shadow: 1px 1px 10px 0px rgba(9, 9, 9, 0.5);
      border-radius: 4px;
    }
  }
  .modal-background {
    position: fixed;
    top: 0;
    left: 0;
    background: rgba(1, 2, 2, 0.5);
    z-index: 999;
    width: 100vw;
    height: 100vh;
  }
}
</style>
