<template>
<!-- <v-container fluid grid-list-xs > -->
  <!-- <v-content> -->
<v-layout row>
    <v-flex xs12>
      <v-card>
        <v-list two-line>
          <v-subheader>Acara dinten niki</v-subheader>
          <template v-for="(item, index) in items">
            
            <v-list-tile
              avatar
              ripple
              @click="toggle(index)"
              :key="item.title"
            >
           
              <v-list-tile-content>
                <v-list-tile-title>{{ item.title }}</v-list-tile-title>
                <v-list-tile-sub-title class="text--primary">{{ item.headline }}</v-list-tile-sub-title>
                <v-list-tile-sub-title>{{ item.subtitle }}</v-list-tile-sub-title>
              </v-list-tile-content>
              <v-list-tile-action>
                <v-list-tile-action-text>{{ item.action }}</v-list-tile-action-text>
                <v-icon color="grey lighten-1">{{ item.status }}</v-icon>
              </v-list-tile-action>
            </v-list-tile>
            <v-divider v-if="index + 1 < items.length" :key="index"></v-divider>
          </template>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
  <!-- </v-content> -->
<!-- </v-container> -->
</template>

<script>
export default {
  name: "TimeList",
  data() {
    return {
      selected: [2],
        items: [
          { action: '15 min', headline: '04:30', title: 'tangi', subtitle: "urip karepe dewe..", status: "alarm_on" },
          { action: '10 min', headline: '04:35', title: 'adus', subtitle: "ben ra lebus..", status: "alarm_on"},
          { action: '1 hr', headline: '06:00', title: 'mangan', subtitle: "urip karepe dewe..", status: "alarm_on"},
          { action: '3 hr', headline: '07:00', title: 'ngoding', subtitle: 'urip karepe dewe..?', status: "alarm_off" },
          { action: '1 hr', headline: '11:00', title: 'turu awan', subtitle: 'urip karepe dewe..?', status: "alarm_on" },
          { action: '30 min', headline: '12:00', title: 'mangan awan', subtitle: 'sego pecel tumpang lawuh ndog bebek karo iwak peyek.', status: "alarm_off" }
        ]
    };
  },
  methods: {
      toggle (index) {
        const i = this.selected.indexOf(index)

        if (i > -1) {
          this.selected.splice(i, 1)
        } else {
          this.selected.push(index)
        }
      }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

