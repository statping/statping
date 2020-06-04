<template>
    <div class="alert alert-warning pb-4 pt-3 mt-5 mb-5" role="alert">
        <h3 class="mb-3">{{message.title}}</h3>
        <span class="mb-3">{{message.description}}</span>
        <div class="row d-block mt-3">
            <span class="col-12 col-md-6 text-left small">
                Started {{niceDate(message.start_on)}} ({{ago(message.start_on)}} ago)
            </span>
            <span class="col-12 col-md-6 text-right float-right small">
                Ends on {{niceDate(message.end_on)}} (in {{ago(message.end_on)}})</span>
        </div>
    </div>
</template>

<script>
export default {
  name: 'MessageBlock',
    props: {
        message: {
            type: Object
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
