<template>
  <li
    v-if="item"
  >
    <button
      class="toggle btn btn-sm d-flex align-items-stretch px-0"
      type="button"
      @click="handleIsOpen"
    >
      <div
        v-if="item.children.length > 0"
        class="toggle-icon pl-2 pr-2 border border-dark rounded-left bg-white d-flex align-items-center"
      >
        <FontAwesomeIcon
          v-if="!isOpen"
          icon="plus-circle"
          class="fa-lg"
        />
        <FontAwesomeIcon
          v-if="isOpen"
          icon="minus-circle"
          class="fa-lg"
        />
      </div>
      
      <div
        v-tooltip="{ content: `${item.parent.downtime ? `${niceDateWithYear(item.parent.downtime.start)} - ${niceDateWithYear(item.parent.downtime.end)}`: ''}`, offset: 5, autoHide: true}"
        class="parent-name rounded-right"
        :class="[{ 'cursor-text rounded': item.children.length === 0 }, item.parent.downtime ? item.parent.downtime.sub_status === 'down' ? 'bg-shade-danger' : 'bg-shade-warning' : 'bg-shade-success']"
      >
        {{ item.parent.name }}
      </div>
    </button>
    
    <ul
      v-if="item.children.length > 0 && isOpen"
      class="list-child pl-0"
    >
      <TreeItem
        v-for="service in item.children"
        :key="service.id"
        class="item"
        :item="service"
      />
    </ul>
  </li>
</template>

<script>
export default {
    name: 'TreeItem',
    props: {
        item: {
            type: Object,
            default: () => null
        }
    },
    data: function () {
        return {
            isOpen: false,
        };
    },
    methods: {
        handleIsOpen: function () {
            this.isOpen = !this.isOpen;
        }
    }
};
</script>

<style scoped>
ul, li {
 margin: 0;
 padding: 0;
 list-style: none;
 display: block;
}

.bg-shade-success {
 background-color: #D6E7D4;
}

.bg-shade-danger {
 background-color: #f7cecc;
}

.bg-shade-warning {
 background-color: #FFE6CC;
}

.cursor-text {
 cursor: text;
}

.toggle {
 line-height: 2 !important;
}

.toggle .toggle-icon {
 border-right: 0 !important;
 z-index: 1;
}

.toggle .parent-name {
 font-size: 16px;
 padding: 0px 5px;
 border: 1px solid;
}

.toggle:focus {
 box-shadow: none;
}

.list-child {
 margin-left: 13px;
 position: relative;
 min-width: 195px;
}

.list-child li {
 padding: 0 0 0 28px;
}

.list-child li::before {
 content: "";
 position: absolute;
 top: 0px;
 left: 5px;
 border-left: 1px solid #A9A9A9;
 border-bottom: 1px solid #A9A9A9;
 width: 15px;
 height: 22px;
 border-radius: 0 0 0 0.3em;
}

.list-child li:not(:last-child)::after{
 position: absolute;
 content: "";
 top: 5px;
 left: 5px;
 border-left: 1px solid #A9A9A9;
 width: 20px;
 height: 100%;
}
</style>