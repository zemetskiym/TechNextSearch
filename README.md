Thank you for your interest in this position.If you are still interested in the role, we request that you complete the below task. 

Please let us know if you have any questions. 
Please don't spend any money to buy/install products/apps to finish this assignment.
Use software tools as you see fit.

Total: 4 questions

1. Backend + Frontend Implementation: 
   Given the raw tabular data at:
   https://www.dropbox.com/scl/fo/c596i84xd39h3tnj283no/h?rlkey=q5k5do6v6yqyyswh2aeyhd61u&dl=0

   use any framework to access/query/update/delete the data and implement a web-based UI interface like Google search. i.e. 

   (a) The UI shall take in the input keyword(s), query the backend database
       to get the data corresponding to the input and display the *matched* results/data in a user-friendly manner.
   
   Details:
      (1) The user must be able to change the input keyword(s) and be able to visualize the results obtained.

      (2) The search should be *fast* returning results within ~5ms.
   
   
   Provide a link to your web-based UI wherever it is hosted.


2. Provide a REST endpoint to the search above 
  
   e.g. A sample REST endpoint: https://<HOST>/api/?query=laser
  
   returns all records that have the keyword "laser" in them. The REST endpoint should output data in JSON format. 
   This should be fast as well.


3. Data Visualization:
   Implement a web-based bar-plot and a line plot (superimposed on the same figure)
   based on "Phase" data column in the given tabular data. i.e. aggregate (simple count will do) 
   on the "Phase" data and plot the results of aggregation of "Phase" column


4. Backend Knowledge:
   Write a *raw* Database query statement (e.g. PostgresSQL: SELECT * FROM table) 
   for "grouping" the records(data) by the "date" field, while aggregating the count of other fields 
   Dump the into a CSV file. Show your work and the database commands used and screenshots.


Assignment Requirements:
=======================
1. The search tool UI should be accessible through a link on the internet.
2. The plot UI should also be accessible through a link on the internet.
3. UI response time is important.
4. For Question (4) show the steps taken to create Database and
   show the operations used to extract required information. This will be asked
   if you are short-listed for the next round.
5. Implementation can be in any stack and hosting system.


Design Guidelines:
=================
The key idea is to communicate backend data in a visually effective manner to
the user. Clean interface, functionality, testability & re-use is preferred.
1. There should be NO hard coding.
2. You can add extra features as you please, if you have time.


Submission Guidelines:
=====================
1. Hosted link(s) for demo showcasing your work
2. Link to the location you have hosted the source code.

Please try to submit the assignment witihin 2 days. 
Note: Time of submission will be noted & so earlier the better!
Note: If you are familiar with any UI/DB stack, this assignment should not take long.

This assignment relates loosely to daily work you will be doing at TechNext
and therefore sets realistic expectations for you and for you to assess us as well.


===================

Details about the data:
The data given has 4 columns:
idx,             <= index
idx_2,           <= index repeated (data is not clean as seen from this redundancy)
patent_id,       <= unique ID of patent
patent_text,     <= Raw text explaining what the patent is about
phase,           <= A string that denotes the "phase" category, the project/patent is in,
Date             <= Date of the patent

Data is not squeaky clean and depending on the requirements, you may have to clean the data.

===================


Thanks for your hardwork & patience.

TechNext Inc.

