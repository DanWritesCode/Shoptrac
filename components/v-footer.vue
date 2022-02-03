<template>
  <div>
    <footer id="footer" class="text-center text-md-left">
      <div class="container mt-5 mb-5">
        <div class="row">
          <div class="col-lg-4 col-md-6">
            <logo :hover="true"></logo>
            <h1 class="ff-poppins fs-24 my-3 hover-footer">
              Brand Name
            </h1>
            <p class="text-grey lh-2 fs-16">
              Description
            </p>
          </div>
          <div class="col-lg-2 d-none d-lg-block"></div>
          <div class="col-lg-2 col-md-3">
            <ul class="list-unstyled">
              <li class="fs-18 ff-poppins text-grey text-uppercase">Quick Links</li>
              <li>
                <n-link to="/option-1">Option 1</n-link>
              </li>
              <li>
                <n-link to="/option-2">Option 2</n-link>
              </li>
              <li>
                <n-link to="/option-3">Option 3</n-link>
              </li>
            </ul>
          </div>
          <div class="col-lg-2 col-md-3">
            <ul class="list-unstyled">
              <li class="fs-18 ff-poppins text-grey text-uppercase">
                You should read...
              </li>
                            <li>
                <n-link to="/faq">Frequently Asked Questions</n-link>
              </li>
              <li>
                <n-link to="/terms-of-service">Terms of Service</n-link>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
    <footer id="footer-bottom">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-6 text-left">
            <div>
              <p class="m-0 text-grey-alt">
                &copy;️ 2022 Shoptrac
              </p>
            </div>
          </div>

        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import Logo from './Logo'
export default {
  name: 'VFooter',
  components: { Logo }
}
</script>

<style lang="scss" scoped>
@import '../assets/sass/variables';
@import '../assets/sass/components/core/footer';
</style>
