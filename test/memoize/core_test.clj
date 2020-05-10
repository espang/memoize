(ns memoize.core-test
  (:require [clojure.test :refer [testing is deftest]]
            [memoize.core :as sut]))

(deftest memoize
  (let [f (fn [value]
            (Thread/sleep 500)
            (* 15 value))]
    (testing "map version of memoize"
      (let [mf (sut/memoized f)]
        (is (= (f 10) (mf 10)))
        (is (= (f 10) (mf 10)))))
    (testing "fifo version of memoize"
      (let [mf (sut/memoized f)]
        (is (= (f 10) (mf 10)))
        (is (= (f 10) (mf 10)))))))
