# frozen_string_literal: true
# SPDX-License-Identifier: BSD-3-Clause
require "csv"
require_relative "_harness"
def build_csv
  (0...200).map { |i| [i, "name-#{i}", i * 1, i * 2, i * 3, i * 4].join(",") }.join("\n") + "\n"
end
data = build_csv
rows = CSV.parse(data)
bench("parse-200x6",    500) { CSV.parse(data) }
bench("generate-200x6", 500) { CSV.generate { |c| rows.each { |r| c << r } } }
