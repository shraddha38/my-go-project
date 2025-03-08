import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 50, // 50 virtual users
  duration: '10s', // Test duration
};

export default function () {

//   let responseFromRaceCondition = http.get('http://localhost:8080/incrementWithRaceCondition');
//   // Check if the responses are of status is 200.
//   check(responseFromRaceCondition, {
//     'status is 200': (r) => r.status === 200,
//   });

  // PESSIMISTICALLY
//   let responseFromPessimisticConcurrency = http.get('http://localhost:8080/incrementWithPessimisticLock');
//   check(responseFromPessimisticConcurrency, {
//     'status is 200': (r) => r.status === 200,
//   });

  // OPTIMISTICALLY
    let responseFromOptimisticConcurrency = http.get('http://localhost:8080/incrementWithOptimisticVersioning');
    check(responseFromOptimisticConcurrency, {
        'status is 200': (r) => r.status === 200,
    });


  sleep(0.01);
}