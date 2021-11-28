<template>
  <div class="bg-dark text-grey mb-3 d-block d-lg-none">
    <div class="container-fluid">
      <div class="row align-content-center py-3">
        <div class="col-12 text-right">
          <db-sidebar-button></db-sidebar-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import DbSidebarButton from './SidebarButton'
export default {
  name: 'HorizontalNavbar',
  components: { DbSidebarButton }
}
</script>

<style scoped></style>
