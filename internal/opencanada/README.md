# Open Canada Tool

This directory contains source code for the "Open Canada Tool" used to interact with the Government of Canada's "Open Government" portal.

The intent is for this tool to perform as a medium between a frontend and the Government of Canada's "Open Government" portal.

<!--
This Attribution Statement will need to be printed on any website using the tool:
"Contains information licensed under the Open Government Licence – Canada."

Licence: https://open.canada.ca/en/open-government-licence-canada
-->

## CKAN

Canada's "Open Government" portal is powered by [CKAN][ckan], an open-source Data Management System (DMS).

[ckan]: https://ckan.org/

### CKAN DataStore

A CKAN DataStore provides two methods to access data:
- Downloading files ([documentation][ckan-datastore-download-docs])
- Data Application Programming Interface (API) ([documentation][ckan-datastore-api-ref-docs])

Whenever possible, the "Open Canada Tool" will obtain data using the Data API.

[ckan-datastore-download-docs]: https://docs.ckan.org/en/latest/maintaining/datastore.html#downloading-resources
[ckan-datastore-api-ref-docs]: https://docs.ckan.org/en/latest/maintaining/datastore.html#api-reference

#### CKAN terminology

| Data Type | Definition |
|-----------|------------|
| DataStore | The CKAN mechanism that provides the Database and access to it. |
| Database | The entire collection of data. |
| Dataset (aka Package) | A parcel of data in the Database. A Dataset contains metadata and any number of Resources. |
| Resource | Contains data in a particular strucutre within the context of the Dataset and/or in a particular format. A Resource contains Records. |
| Record | A representation of a single element of data. A Record contains Fields. |
| Field | A single characteristic of a Record. The combination of Fields typically make the Record they belong to unique within the Resource. |
