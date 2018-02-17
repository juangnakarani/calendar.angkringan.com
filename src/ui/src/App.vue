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
              <v-checkbox color="green" label="Tangi" v-model="rightcheck.acaraku1" light hide-details></v-checkbox>
            <v-checkbox color="red" label="Tahu Tech" v-model="rightcheck.acaraku2" light hide-details></v-checkbox>
            <v-checkbox color="blue" label="Coding" v-model="rightcheck.acaraku3" light hide-details></v-checkbox>
            <v-checkbox color="yellow" label="Dolan" v-model="rightcheck.acaraku4" light hide-details></v-checkbox>
            </div>
        </v-card>
    </v-expansion-panel-content>
    <v-expansion-panel-content>
      <div class="subheading" slot="header">Acara dia</div>
      <v-card flat>
          <div class="ml-2">
            <v-checkbox color="blue" label="Dolan tipis" v-model="rightcheck.acaradia1" light hide-details></v-checkbox>
            <v-checkbox color="red" label="Dolan tenan" v-model="rightcheck.acaradia2" light hide-details></v-checkbox>
          </div>
        </v-card>
    </v-expansion-panel-content>
  </v-expansion-panel>
   
      
    </v-navigation-drawer>
    <v-toolbar fixed app :clipped-left="clipped" color="blue lighten-1">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title v-text="title"></v-toolbar-title>
      <v-spacer></v-spacer>
      <v-text-field prepend-icon="search" hide-details single-line></v-text-field>
      <v-btn icon @click.stop="rightDrawer = !rightDrawer">
        <v-icon>more_vert</v-icon>
      </v-btn>
       <!-- <v-text>Cah Angkringan</v-text> -->
    </v-toolbar>
    <v-content>
      <!-- <v-container fluid>
        <transition mode="out-in"> -->
      <router-view></router-view>
      <!-- </transition>
      </v-container> -->
    </v-content>
    <v-navigation-drawer
      temporary
      :right="right"
      v-model="rightDrawer"
      fixed
    >
      <v-list two-line>
          <!-- <template> -->
        <v-list-tile>
         <v-list-tile-avatar>
            <img src="/assets/nella.png">
          </v-list-tile-avatar>
          <v-list-tile-content>
          <v-list-tile-title>Telo</v-list-tile-title>
          <v-list-tile-sub-title class="text--primary">cah.angkringan@gmail.com</v-list-tile-sub-title>
        </v-list-tile-content>
        </v-list-tile>
<v-divider></v-divider>

          <!-- </template> -->
      </v-list>
       <v-card>
        <!-- <v-card-media src="/static/doc-images/cards/desert.jpg" height="200px">
        </v-card-media> -->
        <v-card-title primary-title>
          <div>
            <h3 class="headline mb-0">User detail</h3>
            <div>Lorem ipsum dolor sit amet amet <br>Lorem ipsum dolor sit amet amet nuwun sewu</div>
          </div>
        </v-card-title>
        <v-card-actions>
            <v-layout justify-space-around>
          <v-btn flat color="blue" @click="push()">My Account</v-btn>
          <v-btn flat color="blue">Logout</v-btn>
            </v-layout>
        </v-card-actions>
      </v-card>
    </v-navigation-drawer>
    <v-footer :fixed="fixed" app class="pa-3">
      <v-spacer></v-spacer>
      <span>&copy; {{ new Date().getFullYear() }} Cah Angkringan</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      picker: null,
      clipped: true,
      drawer: true,
      fixed: true,
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: "Calendar Angkringan",
      rightcheck: {
        acaraku1: true,
        acaraku2: true,
        acaraku3: true,
        acaraku4: false,
        acaradia1: true,
        acaradia2: true,
      }
    };
  },
  methods: {
    push() {
      this.$router.push("/myaccount");
    },
    gohome(){
      this.$router.push("/");
    }
  },
  watch: {
    picker: function(){
      this.gohome()
    }
  }
};
</script>
