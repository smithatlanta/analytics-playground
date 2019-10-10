# Analytics Test Website

- The idea here is to create a test plan to execute against the analytics platform.  
- The test plan would include the following:
  - The number of users to test with.
  - The number of events each user would send.
  - What type of events each user would send.
- Sitting on top of these items would be the ability to test experiments.
  - My thinking on this...
    - Cohort groups are just a grouping of the above users.
    - Experiments are just a grouping of the above events along with a cohort group.
- This is something that will evolve.  It needs to be flexible and easy to use.

## Test Platform Website Requirements

### User Management

#### Required

- Ability to add / remove users.
- Ability to generate a specific number of users(user1, user2…)

##### Table design

tblUser
userid - identity
username - varchar(20)

### Event Management

#### Required

- Ability to add / remove events.
- Ability to generate a specific number of events(event1, event2…).

#### Optional

- Ability to generate events off of a template(or tracking plan)

##### Table design

tblEvent
eventid - identity
eventname - varchar(20)
template - varchar(255)

### Cohort Management

#### Required

- Ability to create / delete cohort groups.
- Ability to add / remove users / percentage of users to cohort group.

##### Table design

tblCohort
cohortid - identity
cohortname - varchar(20)

tblCohortUser
cohortid
userid

### Experiment Management

#### Required

- Ability to create / delete experiments.
- Ability to add / remove events to an experiment.

##### Table design

tblExperiment
experimentid - identity
experimentname - varchar(20)

tblExperimentCohort
experimentid
cohortid

tblExperimentEvent
experimentid
eventid

### Plan Management

#### Required

- Ability to add / remove plans.
- Plans include the following:
  - Number of users to test with(users could be selected or allow the option of selecting a percentage)
  - What events to send to users(events could be selected or allow the option of selecting a percentage)
  - Number of events to send to above users
  - Whether to send events round robin or randomly
    - Round Robin would be user1 gets sent all the events on the plan then user2 gets sent all the events on the plan…
    - Random would be user1 gets sent event10 then user12 gets sent event25. Some users could get more events this way.

##### Table design

```
tblPlan

planid - identity
planname - varchar(20)
sendeventtype - roundrobin or random
planstate - new, execute, inprogress, or complete
```
```
tblPlanEventUser

planeventuserid - identity
planid - foreign key
eventid - foreign key
userid - foriegn key
eventcount - int
```

## Thoughts

- I am keeping the concept of cohorts and experiments separate from the plan because I think they work better loosely coupled.  This may change.
