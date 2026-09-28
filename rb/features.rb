# CloudflarePages SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module CloudflarePagesFeatures
  def self.make_feature(name)
    case name
    when "base"
      CloudflarePagesBaseFeature.new
    when "test"
      CloudflarePagesTestFeature.new
    else
      CloudflarePagesBaseFeature.new
    end
  end
end
