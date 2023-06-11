package acme

//                                   directory
//                                       |
//                                       +--> newNonce
//                                       |
//           +----------+----------+-----+-----+------------+
//           |          |          |           |            |
//           |          |          |           |            |
//           V          V          V           V            V
//      newAccount   newAuthz   newOrder   revokeCert   keyChange
//           |          |          |
//           |          |          |
//           V          |          V
//        account       |        order --+--> finalize
//                      |          |     |
//                      |          |     +--> cert
//                      |          V
//                      +---> authorization
//                                | ^
//                                | | "up"
//                                V |
//                              challenge
//
//                      ACME Resources and Relationships
