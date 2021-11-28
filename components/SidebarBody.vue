<template>
  <div class="sidebar-body">
    <ul class="sidebar-list">
      <li class="sidebar-title">
        Main Navigation
      </li>
      <li class="sidebar-item">
        <n-link to="/dashboard" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'tachometer-alt']"
              class="mr-3 fa-lg fa-fw"
          />
          Dashboard
        </n-link>
      </li>

      <!-- COGS -->
      <li class="sidebar-item">
        <span @click="toggleDropdown('cogs')"  class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'fighter-jet']"
              class="mr-3 fa-lg fa-fw"
          />
          COGS Config
        </span>
        <transition name="fade">
          <div class="navdrop" v-if="toggledVar['cogs']">
            <div class="sidebar-item">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'search']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Manual Per-Order</span>
              </span>
            </div>
            <div class="sidebar-item">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'calendar-plus']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">CJ</span>
              </span>
            </div>
            <div class="sidebar-item">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'trash']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Oberlo</span>
              </span>
            </div>
            <div class="sidebar-item">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'anchor']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">SIB</span>
              </span>
            </div>
            <div class="sidebar-item">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'sticky-note']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Ali</span>
              </span>
            </div>
            <div class="sidebar-item pb-3">
              <span class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'magic']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Custom API</span>
              </span>
            </div>
          </div>
        </transition>
      </li>
      <!-- END OF COGS TAB -->
      <li class="sidebar-item">
        <n-link to="/discord" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'money-bill-alt']"
              class="mr-3 fa-lg fa-fw"
          />
          Payments
        </n-link>
      </li>
      <li class="sidebar-item">
        <n-link to="/users" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'users']"
              class="mr-3 fa-lg fa-fw"
          />
          Users
        </n-link>
      </li>

      <li class="sidebar-title">
        Admin Tools
      </li>
      <li class="sidebar-item">
        <n-link to="/servers" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'server']"
              class="mr-3 fa-lg fa-fw"
          />
          Servers
        </n-link>
      </li>

      <li class="sidebar-item">
        <n-link to="/balance" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'dollar-sign']"
              class="mr-3 fa-lg fa-fw"
          />
          Finances
        </n-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'SidebarBody',
  data() {
    return {
      toggledVar: {'cogs': false},
    }
  },
  methods: {
    toggleDropdown(id) {
      this.toggledVar[id] = !this.toggledVar[id];
    }
  }
}
</script>

<style lang="scss">
@import '../assets/sass/variables';
.sidebar {
  .navdrop {
    transition: height 300ms ease;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .3s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
  .sidebar-body {
    margin-top: 25px;
    margin-bottom: 25px;
    user-select: none;
    .sidebar-list {
      list-style: none;
      padding: 0;
      height: 100%;
      min-height: 100%;
      li, div {
        padding: 0;
      }
      .sidebar-title {
        margin-left: 30px;
        font-size: 12px;
        font-family: 'Poppins', sans-serif;
        font-weight: 600;
        color: $grey-alt;
        margin-bottom: 20px;
      }
      .sidebar-item {
        margin-bottom: 25px;
        margin-left: 30px;
        .sidebar-link {
          color: $grey;
          font-size: 15px;
          display: block;
          letter-spacing: 0.005em;
          transition: margin-left ease-in-out 150ms;
          &:hover {
            color: $brand;
          }
        }
        .sub {
          margin-top: 10px;
          margin-bottom: -15px;
          margin-left: -10px;
          font-size: 14px;
          .label, .nuxt-link {
            margin-left: -10px;
            font-size: 14px;
          }
          transition-property: width;
          transition-duration: 2s;
          transition-timing-function: linear;
          transition-delay: 1s;
        }
        .nuxt-link-active,
        .nuxt-link-exact-active {
          color: white;
          .left {
            transition: margin-left ease-in-out 150ms;
            border-left: 3px solid $brand;
            background: linear-gradient(
                    90deg,
                    #ef423c -165.38%,
                    rgba(239, 66, 60, 0) 90.74%
            );
            margin-left: -30px;
            width: 20px;
            height: 24px;
            position: absolute;
            display: inline-block;
          }
        }
      }
    }
  }
}
</style>
