/// <reference types="cypress" />

context('Setup Process', () => {

  it('should setup Statping with SQLite', () => {
    cy.visit('/setup', {failOnStatusCode: false})
    cy.get('#db_connection').select('sqlite')
    cy.get('#project').clear().type('Demo Tester')
    cy.get('#description').clear().type('This is a test from Crypress!')
    cy.get('#domain').clear().type('http://localhost:8888')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('#password_confirm').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('#title').should('contain', 'Demo Tester')
    cy.get('#description').should('contain', 'This is a test from Crypress!')
  })

    it('should have sample data', () => {
        cy.visit('/')
      cy.get('#title').should('contain', 'Demo Tester')
      cy.get('#description').should('contain', 'This is a test from Crypress!')
        cy.get('.card').should('have.length', 5)
        cy.get('.group_header').should('have.length', 2)
    })

})
