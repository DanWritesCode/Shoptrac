<template>
  <header>
    <div class="container">
      <div class="row align-items-center align-content-center">
        <div class="col-lg-8">
          <h1 class="ff-poppins-eb text-white mb-4">
            Welcome to Shoptrac
          </h1>
          <div class="row">
            <div class="col-lg-7">
              <p class="lh-2 fs-14 ff-nato-sans font-weight-normal mb-5 mt-2">
                
              </p>
            </div>
          </div>
          <home-banner></home-banner>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import HomeBanner from './HomeBanner'

export default {
  name: 'HomeHeader',
  components: { HomeBanner },
}
</script>
