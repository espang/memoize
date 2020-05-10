(ns memoize.core)

(defn memoized [f]
  (let [state (atom {})]
    (fn [v]
      (if-let [r (get @state v)]
        r
        (let [r (f v)]
          (swap! state assoc v r)
          r)))))

(defn memoized-fifo [f]
  (let [capacity 3
        state    (ref {})
        queue    (ref (clojure.lang.PersistentQueue/EMPTY))]
    (fn [v]
      (if-let [r (get @state v)]
        r
        (let [r (f v)]
          (dosync
           (when (= capacity (count @queue))
             (alter state dissoc (peek @queue))
             (alter queue pop))
           (alter state assoc v r)
           (alter queue conj v)
           r))))))
