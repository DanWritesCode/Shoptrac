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

      <li class="sidebar-item">
        <n-link to="/items" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'sitemap']"
              class="mr-3 fa-lg fa-fw"
          />
          Store Products
        </n-link>
      </li>

      <li class="sidebar-item">
        <n-link to="/orders" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'file-invoice-dollar']"
              class="mr-3 fa-lg fa-fw"
          />
          Order Information
        </n-link>
      </li>

      <li class="sidebar-item">
        <n-link to="/customers" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'user']"
              class="mr-3 fa-lg fa-fw"
          />
          Customer Information
        </n-link>
      </li>

      <li class="sidebar-item">
        <n-link to="/revenue" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'circle-dollar-to-slot']"
              class="mr-3 fa-lg fa-fw"
          />
          All Revenue
        </n-link>
      </li>

      <li class="sidebar-item">
        <n-link to="/expenses" class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'comment-dollar']"
              class="mr-3 fa-lg fa-fw"
          />
          All Expenses
        </n-link>
      </li>

      <!-- COGS -->
      <li class="sidebar-item">
        <span @click="toggleDropdown('conf')"  class="sidebar-link">
          <span class="left"></span>
          <font-awesome-icon
              :icon="['fas', 'gear']"
              class="mr-3 fa-lg fa-fw"
          />
          Configuration
        </span>
        <transition name="fade">
          <div class="navdrop" v-if="toggledVar['conf']">
            <div class="sidebar-item">
              <n-link to="/config/cogs" class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'cubes']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Cost of Goods Sold</span>
              </n-link>
            </div>
            <div class="sidebar-item">
              <n-link to="/config/marketing" class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'rectangle-ad']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Marketing Platforms</span>
              </n-link>
            </div>
            <div class="sidebar-item">
              <n-link to="/config/payment-gateways" class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'credit-card']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Payment Gateways</span>
              </n-link>
            </div>
            <div class="sidebar-item">
              <n-link to="/config/recurring-expenses" class="sidebar-link sub">
                <span class="left"></span>
                <font-awesome-icon
                    :icon="['fas', 'calendar']"
                    class="mr-3 fa-lg fa-fw"
                />
                <span class="label">Recurring Expenses</span>
              </n-link>
            </div>
          </div>
        </transition>
      </li>
      <!-- END OF CONFIG TAB -->

    </ul>
  </div>
</template>

<script>
export default {
  name: 'SidebarBody',
  data() {
    return {
      toggledVar: {'conf': false},
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
