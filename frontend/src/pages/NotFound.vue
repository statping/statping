<template>
    <div class="row mt-4">
        <div class="mx-auto">
            <img alt="Statping 404" class="" style="max-width:480px" src="banner.png">
        </div>

        <div class="col-12 mt-5 mb-5">
            <div class="text-center">
                <h2>Page Not Found</h2>
                <h5 class="text-muted">This URL doesn't seem to be available</h5>

                <router-link class="btn btn-outline-success mt-4" to="/">Back To Homepage</router-link>
            </div>
        </div>

    </div>
</template>

<script>
export default {
  name: 'NotFound',
}
</script>
