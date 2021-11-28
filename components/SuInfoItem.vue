<template>
  <div class="row align-items-center align-content-center py-2">
    <div class="col-4 col-lg-3">
      <p
        :class="currentColor + '-text'"
        class="d-inline font-weight-bold fs-14"
      >
        {{ name }}
      </p>
    </div>
    <div class="col-4 d-none d-bp d-xl-block col-lg-4">
      <p class="d-inline">Test</p>
    </div>
    <div class="col-md-2 col-lg-3 d-none d-md-flex">
      <span :class="currentColor" class="status">
        {{ statusMessage }}
      </span>
    </div>
    <div class="col-1 col-lg-2 offset-4 offset-md-4 offset-lg-1 offset-xl-0">
      <button
        @click="toggleDropdown()"
        class="mr-3 btn btn-block btn-custom"
      >View</button>
      <transition name="fade" style="z-index: 1000;">
        <dropdown v-if="isToggled" style="z-index: 1000;">
          <ul class="list-unstyled">
            <li>
              <a class="cursor">
                Option 1
              </a>
            </li>
            <li>
              <a class="cursor">
                Option 2
              </a>
            </li>
            <li>
              <a class="cursor">
                Option 3
              </a>
            </li>
            <li>
              <a class="cursor">
                Option 4
              </a>
            </li>
          </ul>
        </dropdown>
      </transition>
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
  computed: {
    checkIfPending() {
      if (this.statusMessage === 'Pending' || this.statusMessage === 'Success') {
        return 'text-white text-red-hover cursor'
      }
      return 'text-grey-dark'
    }
  },
  mounted() {
    this.currentStatus()
    window.addEventListener('click', this.close)
    window.$nuxt.$on('closeModal', (e) => {
      // handle modal closes here
    })
  },
  beforeDestroy() {
    window.removeEventListener('click', this.close)
  },
  methods: {
    currentStatus() {
      switch (this.status) {
        case 'SUCCESS':
          this.currentColor = 'is-green'
          this.statusMessage = 'Success'
          break
        case 'INVALID':
          this.currentColor = 'is-yellow'
          this.statusMessage = 'Invalid'
          break
        case 'Pending Response':
          this.currentColor = 'is-blue'
          this.statusMessage = 'Replied'
          break
        case 'FAILED':
          this.currentColor = 'is-red'
          this.statusMessage = 'Failed'
          break
        case 'FAIL':
          this.currentColor = 'is-red'
          this.statusMessage = 'Failed'
          break
        case 'UNKNOWN':
          this.currentColor = 'is-red'
          this.statusMessage = 'Unknown'
          break
        default:
          this.currentColor = 'is-blue'
          this.statusMessage = 'Pending'
          break
      }
    },
    toggleDropdown() {
      this.isToggled = !this.isToggled
    },
    close(e) {
      if (!this.$el.contains(e.target)) {
        this.isToggled = false
      }
    },
  }
}
</script>
