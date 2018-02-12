<template>
  <v-app>
    <v-navigation-drawer
      :mini-variant="miniVariant"
      :clipped="clipped"
      v-model="drawer"
      app
      fixed
    >
    <v-layout justify-center>
      <!-- <v-flex xs12> -->
        <v-date-picker color="blue lighten-2" v-model="picker" no-title scrollable actions></v-date-picker>
      <!-- </v-flex> -->
      </v-layout>
  
    <v-expansion-panel flat expand>
    <v-expansion-panel-content v-bind:value="1">
      <div class="subheading" slot="header">Acaraku</div>
      <v-card flat>
            <div class="ml-2">
              <v-checkbox color="green" label="Baca Quran" v-model="rightcheck" light hide-details></v-checkbox>
            <v-checkbox color="red" label="Tahu Tech" v-model="rightcheck" light hide-details></v-checkbox>
            <v-checkbox color="blue" label="Coding" v-model="rightcheck" light hide-details></v-checkbox>
            <v-checkbox color="yellow" label="Dolan" v-model="rightcheck" light hide-details></v-checkbox>
            </div>
        </v-card>
    </v-expansion-panel-content>
    <v-expansion-panel-content>
      <div class="subheading" slot="header">Acara dia</div>
      <v-card flat>
          <div class="ml-2">
            <v-checkbox color="blue" label="Dolan tipis" v-model="rightcheck" light hide-details></v-checkbox>
            <v-checkbox color="red" label="Dolan tenan" v-model="rightcheck" light hide-details></v-checkbox>
          </div>
        </v-card>
    </v-expansion-panel-content>
  </v-expansion-panel>
   
      
    </v-navigation-drawer>
    <v-toolbar fixed app :clipped-left="clipped">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title v-text="title"></v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon @click.stop="rightDrawer = !rightDrawer">
        <v-icon>menu</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
    <v-navigation-drawer
      temporary
      :right="right"
      v-model="rightDrawer"
      fixed
    >
      <v-list>
        <v-list-tile @click.native="right = !right">
          <v-list-tile-action>
            <v-icon>compare_arrows</v-icon>
          </v-list-tile-action>
          <v-list-tile-title>Switch drawer (click me)</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-footer :fixed="fixed" app>
      <span>&copy; 2017 Cah Angkringan</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: true,
      drawer: true,
      fixed: true,
      items: [
        {
          icon: "bubble_chart",
          title: "Inspire"
        }
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: "Calendar Angkringan",
      rightcheck: []
    };
  }
};
</script>
