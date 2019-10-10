# Analytics Plan Runner

This will be a much smarter version of the analytics test client that runs the plans created by the website.

The basic flow of the application will be as follows:

1. Application will poll postgres database plan table every x minutes looking for plans in the execute state.

2. Application updates the plan state to inprogress.

3. Application queries the plan.

4. Application determines what users will be sending what events to the analytics platform.
    
    -
    -

## Required

- keeps track of the total number of events created by users - grafana / cloudwatch
- keeps track of the total number of each event created by users - grafana / cloudwatch

## Optional
