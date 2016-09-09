class CreateFonts < ActiveRecord::Migration
  def change
    create_table :fonts do |t|
      t.string :name, null: false
      t.string :download_url, null: false

      t.timestamps null: false
    end
  end
end
