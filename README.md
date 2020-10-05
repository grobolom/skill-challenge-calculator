# [Skill-Challenge-Calculator](https://skillcalc5e.cc)

Skill Challenge probability calculator for 5e (and 4e). Helps you figure out how
likely your party is to succeed on any particular skill challenge, given their
average skill bonuses and the DC of the challenge that you'd like to set.

At the moment, the golang service is not used - I've moved the logic from
the backend directly into the Vue.js frontend, as it was a significant
performance overhead to be making an API call every time for a relatively
simple math calculation, especially one that is light-weight enought to easily
be handled in the UI. However, I did want to create the API initially to
show a POC for how this Vue <-> golang microservice system would work.

## Future Plans

[] Continue code cleanup
[] Add testing
[] Remove golang service entirely
[] Add screenshots of original Excalidraw mocks
[] Gather feedback