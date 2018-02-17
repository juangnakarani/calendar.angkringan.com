<template>
<v-container fluid grid-list-md>
    <v-flex xs12>
  <form>
    <v-text-field
      v-model="name"
      label="Name"
      :counter="10"
      v-validate="'required|max:10'"
      data-vv-name="name"
      required
    ></v-text-field>
    <v-text-field
      v-model="email"
      label="E-mail"
      v-validate="'required|email'"
      data-vv-name="email"
      required
    ></v-text-field>
    <v-select
      :items="items"
      v-model="select"
      label="Select"
      v-validate="'required'"
      data-vv-name="select"
      required
    ></v-select>
    <v-checkbox
      v-model="checkbox"
      value="1"
      label="Option"
      v-validate="'required'"
      data-vv-name="checkbox"
      type="checkbox"
      required
    ></v-checkbox>

    <v-btn @click="submit">OK</v-btn>
    <v-btn @click="clear">Cancel</v-btn>
  </form>
    </v-flex>
</v-container>
</template>
<script>
export default {
  $_veeValidate: {
      validator: 'new'
    },

    data: () => ({
      name: 'Telo',
      email: 'cah.angkringan@gmail.com',
      select: 'Warga',
      items: [
        'Admin',
        'Warga',
      ],
      checkbox: null,
      dictionary: {
        attributes: {
          email: 'E-mail Address'
          // custom attributes
        },
        custom: {
          name: {
            required: () => 'Name can not be empty',
            max: 'The name field may not be greater than 10 characters'
            // custom messages
          },
          select: {
             required: 'Select field is required'
          }
        }
      }
    }),

    // mounted () {
    //   this.$validator.localize('en', this.dictionary)
    // },

    methods: {
      submit () {
        this.$validator.validateAll()
      },
      clear () {
        this.name = ''
        this.email = ''
        this.select = null
        this.checkbox = null
        this.$validator.reset()
      }
    }
}
</script>
