<template>
  <div class="project-calendar">
    <v-sheet height="64">
        <v-toolbar
          flat
        >
          <v-btn
            outlined
            class="mr-4"
            color="grey darken-2"
            @click="setToday"
          >
            Today
          </v-btn>
          <v-btn
            fab
            text
            small
            color="grey darken-2"
            @click="prev"
          >
            <v-icon small>
              mdi-chevron-left
            </v-icon>
          </v-btn>
          <v-btn
            fab
            text
            small
            color="grey darken-2"
            @click="next"
          >
            <v-icon small>
              mdi-chevron-right
            </v-icon>
          </v-btn>
          <v-toolbar-title v-if="$refs.calendar">
            {{ $refs.calendar.title }}
          </v-toolbar-title>
          <v-spacer></v-spacer>
          <v-menu
            bottom
            right
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                outlined
                color="grey darken-2"
                v-bind="attrs"
                v-on="on"
              >
                <span>{{ typeToLabel[type] }}</span>
                <v-icon right>
                  mdi-menu-down
                </v-icon>
              </v-btn>
            </template>
            <v-list>
              <v-list-item @click="type = 'day'">
                <v-list-item-title>Day</v-list-item-title>
              </v-list-item>
              <v-list-item @click="type = 'week'">
                <v-list-item-title>Week</v-list-item-title>
              </v-list-item>
              <v-list-item @click="type = 'month'">
                <v-list-item-title>Month</v-list-item-title>
              </v-list-item>
              <v-list-item @click="type = '4day'">
                <v-list-item-title>4 days</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </v-toolbar>
      </v-sheet>
      <v-sheet height="600">
        <v-calendar
          ref="calendar"
          v-model="focus"
          color="primary"
          :events="events"
          :event-color="getEventColor"
          :type="type"
          @click:event="showEvent"
          @click:more="viewDay"
          @click:date="viewDay"
          @change="updateRange"
        ></v-calendar>
        <v-menu
          v-model="selectedOpen"
          :close-on-content-click="false"
          :activator="selectedElement"
          offset-x
        >
          <v-card
            color="grey lighten-4"
            min-width="350px"
            flat
          >
            <v-toolbar
              :color="selectedEvent.color"
              dark
            >
              <v-btn icon>
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
              <v-toolbar-title v-html="selectedEvent.name"></v-toolbar-title>
              <v-spacer></v-spacer>
              <v-btn icon>
                <v-icon>mdi-heart</v-icon>
              </v-btn>
              <v-btn icon>
                <v-icon>mdi-dots-vertical</v-icon>
              </v-btn>
            </v-toolbar>
            <v-card-text>
              <span v-html="selectedEvent.details"></span>
            </v-card-text>
            <v-card-actions>
              <v-btn
                text
                color="secondary"
                @click="selectedOpen = false"
              >
                Cancel
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-menu>
      </v-sheet>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: "ProjectCalendar",
  props: ['slice'],
  data: () => ({
      focus: '',
      type: 'month',
      typeToLabel: {
        month: 'Month',
        week: 'Week',
        day: 'Day',
        '4day': '4 Days',
      },
      selectedEvent: {},
      selectedElement: null,
      selectedOpen: false,
      events: [],
      //colors: ['blue', 'indigo', 'deep-purple', 'cyan', 'green', 'orange'],//in order of Roofing, Gutters, Siding, Flooring, Drywall, Painting,
      colors:[
        {key: 'blue', value:'Roofing'},
        {key:'indigo', value: 'Gutters'},
        {key: 'deep-purple', value: 'Siding'},
        {key: 'cyan', value: 'Flooring'},
        {key: 'green', value: 'Drywall'},
        {key: 'orange', value: 'Painting'}
      ],
      names: ['Meeting', 'Holiday', 'PTO', 'Travel', 'Event', 'Birthday', 'Conference', 'Party'],
    }),
    computed: {
      //...mapGetters(['project/projects']),
      projects() {
        var projectsArr = this.$store.state.project.projects
        for(var x = 0; x < projectsArr.length; x++) {
          for (var y = 0; y < this.colors.length; y++) {
            if (projectsArr[x].type == this.colors[y].value) {
              projectsArr[x].color = this.colors[y].key
            }
          }
        }
        return projectsArr
      }
    },
    mounted () {
      this.$refs.calendar.checkChange()
    },
    methods: {
      viewDay ({ date }) {
        this.focus = date
        this.type = 'day'
      },
      getEventColor (event) {
        return event.color
      },
      setToday () {
        this.focus = ''
      },
      prev () {
        this.$refs.calendar.prev()
      },
      next () {
        this.$refs.calendar.next()
      },
      showEvent ({ nativeEvent, event }) {
        const open = () => {
          this.selectedEvent = event
          this.selectedElement = nativeEvent.target
          setTimeout(() => {
            this.selectedOpen = true
          }, 10)
        }

        if (this.selectedOpen) {
          this.selectedOpen = false
          setTimeout(open, 10)
        } else {
          open()
        }

        nativeEvent.stopPropagation()
      },
      updateRange ({ start, end }) {
        const events = []

        const min = new Date(`${start.date}T00:00:00`)
        const max = new Date(`${end.date}T23:59:59`)
        const days = (max.getTime() - min.getTime()) / 86400000
        const eventCount = this.$store.state.project.projects ? this.$store.state.project.projects.length : null

        for (let i = 0; i < eventCount; i++) {
          // const allDay = this.rnd(0, 3) === 0
          // const firstTimestamp = this.rnd(min.getTime(), max.getTime())
          // const first = new Date(firstTimestamp - (firstTimestamp % 900000))
          // const secondTimestamp = this.rnd(2, allDay ? 288 : 8) * 900000
          // const second = new Date(first.getTime() + secondTimestamp)
          const startTimestamp = this.$store.state.project.projects[i].start.seconds
          const endTimestamp = this.$store.state.project.projects[i].end.seconds
          const start = new Date(startTimestamp * 1000)
          const end = new Date(endTimestamp * 1000)

          events.push({
            name: this.projects[i].name,
            color: this.projects[i].color,
            client: this.projects[i].client,
            type: this.projects[i].type,
            start: start,
            end: end,
            //timed: !allDay,
          })

          
        }

        this.events = events
      },
      rnd (a, b) {
        return Math.floor((b - a + 1) * Math.random()) + a
      },
    },
}
</script>
<style lang="scss">
.project-calendar {
  max-width:140rem;
  width:100%;
  margin:auto;
}
</style>