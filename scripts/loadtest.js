import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 10, // 10 virtual users
  duration: '5s', // Test duration
};

export default function () {

  // RACE CONDITION
  // let responseFromRaceCondition = http.get('http://localhost:8080/incrementWithRaceCondition');
  // check(responseFromRaceCondition, {
  //   'status is 200': (r) => r.status === 200,
  // });

  // PESSIMISTICALLY
  // let responseFromPessimisticConcurrency = http.get('http://localhost:8080/incrementWithPessimisticLock');
  // check(responseFromPessimisticConcurrency, {
  //   'status is 200': (r) => r.status === 200,
  // });

  // OPTIMISTICALLY
    let responseFromOptimisticConcurrency = http.get('http://localhost:8080/incrementWithOptimisticVersioning');
    check(responseFromOptimisticConcurrency, {
        'status is 200': (r) => r.status === 200,
    });


  sleep(0.01);
}