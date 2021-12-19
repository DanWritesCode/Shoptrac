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
          <div class="text-center mt-4" v-if="!installed && !error"><h1>Installing StonksUp...</h1></div>
          <div class="text-center mt-4" v-if="installed"><h1>Thank you for installing StonksUp!</h1></div>
          <div class="text-center mt-4 text-red" v-if="errorMsg"><h3>{{errorMsg}}</h3></div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>

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
      installed: false,
      error: false,
      errorMsg: ""
    }
  },
  mounted() {
    let urlParams = new URLSearchParams(window.location.search);
    const code = decodeURIComponent(urlParams.get("code")).trim();
    const hmac = decodeURIComponent(urlParams.get("hmac")).trim();
    const host = decodeURIComponent(urlParams.get("host")).trim();
    const shop = decodeURIComponent(urlParams.get("shop")).trim();
    const state = decodeURIComponent(urlParams.get("state")).trim();

    this.installAPI(window.location.search, code, hmac, host, shop, state);
  },
  methods: {
    async installAPI(url, code, hmac, host, shop, state) {

      await this.$axios
          .$post("http://localhost:8000/install", {
            "url": url, "code": code, "hmac": hmac, "host": host, "shop": shop, "nonce": state
          })
          .then((res) => {
            this.installed = true
          })
          .catch((res) => {
            if(res !== undefined)
              this.errorMsg = res;
            this.error = true;
          });
    }
  },

}
</script>
