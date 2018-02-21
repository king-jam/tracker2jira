* Notes:
 * Activity Endpoint PivotalTracker only provides the last 6 months of data
 * Need a validation step on adding a project with a field to determine
 * if a manual update based on iterating all types is necessary before
 * moving to an activity based delta iterator

 * only sync events after current timestamp of project creation to avoid losing
   activity stream
