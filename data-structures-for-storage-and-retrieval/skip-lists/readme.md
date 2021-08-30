# Skip Lists

## Notes

### LeveDB / embedded K/V store

- What is the API / what do we want to acomplish?
  - get(key: string)
  - set (key: string, value: string)
  - delete(key string)
  - range scan()
    - iterate through a range of keys(e.g keys between "A" and "C", keys starting with some prefix)
- What propertiers do we want out of this system?
  - how will clients access this system?
    - clients are just gonna import some library
    - not separate server for the DB, rather embedded in the application
      - PRO: avoid a lot of network traffic
      - CONS: all clients have to be in one machine -> files are local??
        - COUPLING
        - Single Point of Failure
        - No redundancy
  - that library will directly access files on disk
  - parallel access / high level of cuncurrency
- what sort of "reliability" / "durability" / "persistence" / "consistency" guarantees do we want?
  - maybe configurable?
  - "we want our writes to be durable"
    - if set operatino succeeds -> we should not lose our data at that point?
  
*Design in Notes*