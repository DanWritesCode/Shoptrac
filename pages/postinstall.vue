<template>
  <div class="h-100">
    <div class="row justify-content-center align-content-center h-100">
      <div class="col-10 col-md-5 mb-5">
        <div class="card bg-box p-5 text-center">
          <div class="mb-3 text-center"><img
              src="~/assets/images/logo.png"
              alt="StonksUp"
              class="logo"
          />
          </div>
          <div class="text-center mt-4" :class="txtColor"><h1>{{msg}}</h1></div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>

import {API_URL} from "@/config";

export default {
  name: 'PostInstall',
  components: {},
  layout: 'default',
  transition: 'fade',
  head() {
    return {
      title: 'StonksUp | App Installation'
    }
  },
  data() {
    return {
      msg: "Installing StonksUp...",
      txtColor: "text-white"
    }
  },
  mounted() {
    // https://github.com/nuxt/nuxt.js/issues/5703#issuecomment-563164340
    // when you have a 2-year-old bug in Nuxt...
    if (this.$nuxt.layoutName === 'default') {
      let urlParams = new URLSearchParams(window.location.search);
      const code = decodeURIComponent(urlParams.get("code")).trim();
      const hmac = decodeURIComponent(urlParams.get("hmac")).trim();
      const host = decodeURIComponent(urlParams.get("host")).trim();
      const shop = decodeURIComponent(urlParams.get("shop")).trim();
      const state = decodeURIComponent(urlParams.get("state")).trim();

      this.$axios
          .$post(`${API_URL}/install`, {
            "url": window.location.search, "code": code, "hmac": hmac, "host": host, "shop": shop, "nonce": state
          })
          .then((res) => {
            this.msg = "Thank you for installing StonksUp!";
          })
          .catch((res) => {
            this.msg = "An error occurred while installing the app! " + res
            this.txtColor = "text-red"
          });
    }
  },

}
</script>
